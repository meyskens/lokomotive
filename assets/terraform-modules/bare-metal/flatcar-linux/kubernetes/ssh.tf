# Secure copy etcd TLS assets and kubeconfig to controllers. Activates kubelet.service
resource "null_resource" "copy-controller-secrets" {
  count = length(var.controller_names)

  # Without depends_on, remote-exec could start and wait for machines before
  # matchbox groups are written, causing a deadlock.
  depends_on = [
    matchbox_group.install,
    matchbox_group.controller,
    matchbox_group.worker,
  ]

  connection {
    type    = "ssh"
    host    = var.controller_domains[count.index]
    user    = "core"
    timeout = "60m"
  }

  provisioner "file" {
    content = var.enable_tls_bootstrap ? templatefile("${path.module}/cl/bootstrap-kubeconfig.yaml.tmpl", {
      token_id     = random_string.bootstrap_token_id_controller[0].result
      token_secret = random_string.bootstrap_token_secret_controller[0].result
      ca_cert      = module.bootkube.ca_cert
      server       = "https://${var.k8s_domain_name}:6443"
    }) : module.bootkube.kubeconfig-kubelet
    destination = "$HOME/kubeconfig"
  }

  provisioner "file" {
    content     = module.bootkube.etcd_ca_cert
    destination = "$HOME/etcd-client-ca.crt"
  }

  provisioner "file" {
    content     = module.bootkube.etcd_client_cert
    destination = "$HOME/etcd-client.crt"
  }

  provisioner "file" {
    content     = module.bootkube.etcd_client_key
    destination = "$HOME/etcd-client.key"
  }

  provisioner "file" {
    content     = module.bootkube.etcd_server_cert
    destination = "$HOME/etcd-server.crt"
  }

  provisioner "file" {
    content     = module.bootkube.etcd_server_key
    destination = "$HOME/etcd-server.key"
  }

  provisioner "file" {
    content     = module.bootkube.etcd_peer_cert
    destination = "$HOME/etcd-peer.crt"
  }

  provisioner "file" {
    content     = module.bootkube.etcd_peer_key
    destination = "$HOME/etcd-peer.key"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo mkdir -p /etc/ssl/etcd/etcd",
      "sudo mv etcd-client* /etc/ssl/etcd/",
      "sudo cp /etc/ssl/etcd/etcd-client-ca.crt /etc/ssl/etcd/etcd/server-ca.crt",
      "sudo mv etcd-server.crt /etc/ssl/etcd/etcd/server.crt",
      "sudo mv etcd-server.key /etc/ssl/etcd/etcd/server.key",
      "sudo cp /etc/ssl/etcd/etcd-client-ca.crt /etc/ssl/etcd/etcd/peer-ca.crt",
      "sudo mv etcd-peer.crt /etc/ssl/etcd/etcd/peer.crt",
      "sudo mv etcd-peer.key /etc/ssl/etcd/etcd/peer.key",
      "sudo chown -R etcd:etcd /etc/ssl/etcd",
      "sudo chmod -R 500 /etc/ssl/etcd",
      "sudo mv $HOME/kubeconfig /etc/kubernetes/kubeconfig",
    ]
  }
}

# Secure copy kubeconfig to all workers. Activates kubelet.service
resource "null_resource" "copy-worker-secrets" {
  count = length(var.worker_names)

  # Without depends_on, remote-exec could start and wait for machines before
  # matchbox groups are written, causing a deadlock.
  depends_on = [
    matchbox_group.install,
    matchbox_group.controller,
    matchbox_group.worker,
  ]

  connection {
    type    = "ssh"
    host    = var.worker_domains[count.index]
    user    = "core"
    timeout = "60m"
  }

  provisioner "file" {
    content = var.enable_tls_bootstrap ? templatefile("${path.module}/cl/bootstrap-kubeconfig.yaml.tmpl", {
      token_id     = random_string.bootstrap_token_id_worker[0].result
      token_secret = random_string.bootstrap_token_secret_worker[0].result
      ca_cert      = module.bootkube.ca_cert
      server       = "https://${var.k8s_domain_name}:6443"
    }) : module.bootkube.kubeconfig-kubelet
    destination = "$HOME/kubeconfig"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo mv $HOME/kubeconfig /etc/kubernetes/kubeconfig",
    ]
  }

  # Triggered when the Ignition Config changes
  triggers = {
    ignition_config = null_resource.reprovision-worker-when-ignition-changes[count.index].id
  }

}

# Secure copy bootkube assets to ONE controller and start bootkube to perform
# one-time self-hosted cluster bootstrapping.
resource "null_resource" "bootkube-start" {
  # Without depends_on, this remote-exec may start before the kubeconfig copy.
  # Terraform only does one task at a time, so it would try to bootstrap
  # while no Kubelets are running.
  depends_on = [
    null_resource.copy-controller-secrets,
    null_resource.copy-worker-secrets,
  ]

  connection {
    type    = "ssh"
    host    = var.controller_domains[0]
    user    = "core"
    timeout = "15m"
  }

  provisioner "file" {
    source      = var.asset_dir
    destination = "$HOME/assets"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo mv $HOME/assets /opt/bootkube",
      # Use stdbuf to disable the buffer while printing logs to make sure everything is transmitted back to
      # Terraform before we return error. We should be able to remove it once
      # https://github.com/hashicorp/terraform/issues/27121 is resolved.
      "sudo systemctl start bootkube || (stdbuf -i0 -o0 -e0 sudo journalctl -u bootkube --no-pager; exit 1)",
    ]
  }
}

resource "null_resource" "reprovision-controller-when-ignition-changes" {
  count = length(var.controller_names)
  # Triggered when the Ignition Config changes
  triggers = {
    ignition_config = matchbox_profile.controllers[count.index].raw_ignition
  }
  # Wait for the new Ignition config object to be ready before rebooting
  depends_on = [matchbox_group.controller]
  # Trigger running Ignition on the next reboot (first_boot flag file) and reboot the instance, or, if the instance needs to be (re)provisioned, run external commands for PXE booting (also runs on the first provisioning)
  provisioner "local-exec" {
    command = templatefile("${path.module}/pxe-helper.sh.tmpl", { domain = var.controller_domains[count.index], name = var.controller_names[count.index], mac = var.controller_macs[count.index], pxe_commands = var.pxe_commands })
  }
}

resource "null_resource" "reprovision-worker-when-ignition-changes" {
  count = length(var.worker_names)
  # Triggered when the Ignition Config changes
  triggers = {
    ignition_config = matchbox_profile.workers[count.index].raw_ignition
  }
  # Wait for the new Ignition config object to be ready before rebooting
  depends_on = [matchbox_group.worker]
  # Trigger running Ignition on the next reboot (first_boot flag file) and reboot the instance, or, if the instance needs to be (re)provisioned, run external commands for PXE booting (also runs on the first provisioning)
  provisioner "local-exec" {
    command = templatefile("${path.module}/pxe-helper.sh.tmpl", { domain = var.worker_domains[count.index], name = var.worker_names[count.index], mac = var.worker_macs[count.index], pxe_commands = var.pxe_commands })
  }
}
