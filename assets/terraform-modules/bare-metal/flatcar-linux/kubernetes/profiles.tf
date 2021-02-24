// Flatcar Container Linux install profile (from release.flatcar-linux.net)
resource "matchbox_profile" "flatcar-install" {
  count = length(var.controller_names) + length(var.worker_names)
  name = format(
    "%s-flatcar-install-%s",
    var.cluster_name,
    concat(var.controller_names, var.worker_names)[count.index]
  )

  kernel = "${var.download_protocol}://${var.os_channel}.release.flatcar-linux.net/amd64-usr/${var.os_version}/flatcar_production_pxe.vmlinuz"

  initrd = [
    "${var.download_protocol}://${var.os_channel}.release.flatcar-linux.net/amd64-usr/${var.os_version}/flatcar_production_pxe_image.cpio.gz",
  ]

  args = flatten([
    "initrd=flatcar_production_pxe_image.cpio.gz",
    "ignition.config.url=${var.matchbox_http_endpoint}/ignition?uuid=$${uuid}&mac=$${mac:hexhyp}",
    "flatcar.first_boot=yes",
    "console=tty0",
    "console=ttyS0",
    var.kernel_args,
  ])

  container_linux_config = templatefile("${path.module}/cl/install.yaml.tmpl", {
    os_channel               = var.os_channel
    os_version               = var.os_version
    ignition_endpoint        = format("%s/ignition", var.matchbox_http_endpoint)
    install_disk             = var.install_disk
    container_linux_oem      = var.container_linux_oem
    ssh_keys                 = jsonencode(var.ssh_keys)
    install_to_smallest_disk = var.install_to_smallest_disk
    # only cached-container-linux profile adds -b baseurl
    baseurl_flag = ""
  })
}

// Flatcar Container Linux Install profile (from matchbox /assets cache)
// Note: Admin must have downloaded os_version into matchbox assets/flatcar.
resource "matchbox_profile" "cached-flatcar-linux-install" {
  count = length(var.controller_names) + length(var.worker_names)
  name = format(
    "%s-cached-flatcar-linux-install-%s",
    var.cluster_name,
    concat(var.controller_names, var.worker_names)[count.index]
  )

  kernel = "/assets/flatcar/${var.os_version}/flatcar_production_pxe.vmlinuz"

  initrd = [
    "/assets/flatcar/${var.os_version}/flatcar_production_pxe_image.cpio.gz",
  ]

  args = flatten([
    "initrd=flatcar_production_pxe_image.cpio.gz",
    "ignition.config.url=${var.matchbox_http_endpoint}/ignition?uuid=$${uuid}&mac=$${mac:hexhyp}",
    "flatcar.first_boot=yes",
    "console=tty0",
    "console=ttyS0",
    var.kernel_args,
  ])

  container_linux_config = templatefile("${path.module}/cl/install.yaml.tmpl", {
    os_channel               = var.os_channel
    os_version               = var.os_version
    ignition_endpoint        = format("%s/ignition", var.matchbox_http_endpoint)
    install_disk             = var.install_disk
    container_linux_oem      = var.container_linux_oem
    ssh_keys                 = jsonencode(var.ssh_keys)
    install_to_smallest_disk = var.install_to_smallest_disk
    # profile uses -b baseurl to install from matchbox cache
    baseurl_flag = "-b ${var.matchbox_http_endpoint}/assets/flatcar"
  })
}

// Kubernetes Controller profiles
resource "matchbox_profile" "controllers" {
  count = length(var.controller_names)
  name = format(
    "%s-controller-%s",
    var.cluster_name,
    var.controller_names[count.index]
  )
  raw_ignition = module.controller[count.index].clc_config
}

// Kubernetes Worker profiles
resource "matchbox_profile" "workers" {
  count = length(var.worker_names)
  name = format(
    "%s-worker-%s",
    var.cluster_name,
    var.worker_names[count.index]
  )
  raw_ignition = module.worker[count.index].clc_config
}
