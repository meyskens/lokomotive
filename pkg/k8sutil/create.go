// Copyright 2020 The Lokomotive Authors
// Copyright 2019 The Kubernetes Authors
// Copyright 2015 CoreOS, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8sutil

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	yamlserializer "k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/apimachinery/pkg/util/yaml"
	corev1typed "k8s.io/client-go/kubernetes/typed/core/v1"
	sigyaml "sigs.k8s.io/yaml"

	"github.com/kinvolk/lokomotive/internal"
)

// Namespace struct for holding the Lokomotive specific metadata.
// when installing cluster or components.
type Namespace struct {
	Name        string
	Labels      map[string]string
	Annotations map[string]string
}

// Adapted from https://github.com/kubernetes-incubator/bootkube/blob/83d32756c6b02c26cab1de3f03b57f06ae4339a7/pkg/bootkube/create.go

type manifest struct {
	kind       string
	apiVersion string
	namespace  string
	name       string
	raw        []byte

	filepath string
}

func (m manifest) String() string {
	if m.namespace == "" {
		return fmt.Sprintf("%s %s %s", m.filepath, m.kind, m.name)
	}
	return fmt.Sprintf("%s %s %s/%s", m.filepath, m.kind, m.namespace, m.name)
}

func (m manifest) Kind() string {
	return m.kind
}

func (m manifest) Raw() []byte {
	return m.raw
}

func (m manifest) Name() string {
	return m.name
}

// LoadManifests parses a map of Kubernetes manifests.
func LoadManifests(files map[string]string) ([]manifest, error) {
	var manifests []manifest
	for path, fileContent := range files {
		r := strings.NewReader(fileContent)
		ms, err := parseManifests(r)
		if err != nil {
			return nil, fmt.Errorf("parsing manifest %q: %w", path, err)
		}
		manifests = append(manifests, ms...)
	}
	return manifests, nil
}

// parseManifests parses a YAML or JSON document that may contain one or more
// kubernetes resources.
func parseManifests(r io.Reader) ([]manifest, error) {
	reader := yaml.NewYAMLReader(bufio.NewReader(r))
	var manifests []manifest
	for {
		yamlManifest, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				return manifests, nil
			}
			return nil, err
		}
		yamlManifest = bytes.TrimSpace(yamlManifest)
		if len(yamlManifest) == 0 {
			continue
		}

		jsonManifest, err := yaml.ToJSON(yamlManifest)
		if err != nil {
			return nil, fmt.Errorf("invalid manifest: %w", err)
		}
		m, err := parseJSONManifest(jsonManifest)
		if err != nil {
			return nil, fmt.Errorf("parse manifest: %w", err)
		}
		manifests = append(manifests, m...)
	}
}

// parseJSONManifest parses a single JSON Kubernetes resource.
func parseJSONManifest(data []byte) ([]manifest, error) {
	if string(data) == "null" {
		return nil, nil
	}
	var m struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
		} `json:"metadata"`
	}
	if err := json.Unmarshal(data, &m); err != nil {
		return nil, fmt.Errorf("unmarshaling manifest: %w", err)
	}

	// We continue if the object we received was a *List kind. Otherwise if a
	// single object is received we just return from here.
	if !strings.HasSuffix(m.Kind, "List") {
		return []manifest{{
			kind:       m.Kind,
			apiVersion: m.APIVersion,
			namespace:  m.Metadata.Namespace,
			name:       m.Metadata.Name,
			raw:        data,
		}}, nil
	}

	// We parse the list of items and extract one object at a time
	var mList struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Metadata   struct {
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
		} `json:"metadata"`
		Items []json.RawMessage `json:"items"`
	}
	if err := json.Unmarshal(data, &mList); err != nil {
		return nil, fmt.Errorf("unmarshaling manifest list: %w", err)
	}
	var manifests []manifest
	for _, item := range mList.Items {
		// make a recursive call, since this is a single object it will be
		// parsed and returned to us
		mn, err := parseJSONManifest(item)
		if err != nil {
			return nil, err
		}
		manifests = append(manifests, mn...)
	}
	return manifests, nil
}

// CreateOrUpdateNamespace creates the release namespace or updates the namespace
// if it already exists.
func CreateOrUpdateNamespace(ns Namespace, nsclient corev1typed.NamespaceInterface) error {
	if ns.Name == "" {
		return fmt.Errorf("namespace name can't be empty")
	}

	namespace, err := nsclient.Get(context.TODO(), ns.Name, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return createNamespace(ns, nsclient)
		}

		return fmt.Errorf("getting namespace %q: %w", ns.Name, err)
	}
	// Namespace exists, hence updating the namespace.
	return updateNamespace(namespace, ns, nsclient)
}

// updateNamespace updates an existing namespace.
func updateNamespace(namespace *v1.Namespace, ns Namespace, nsclient corev1typed.NamespaceInterface) error {
	// Merge new labels and annotations with existing ones.
	updatedLabels := internal.MergeMaps(ns.Labels, namespace.ObjectMeta.Labels)
	updatedAnnotations := internal.MergeMaps(ns.Annotations, namespace.ObjectMeta.Annotations)

	namespace.ObjectMeta.Labels = updatedLabels
	namespace.ObjectMeta.Annotations = updatedAnnotations

	if _, err := nsclient.Update(context.TODO(), namespace, metav1.UpdateOptions{}); err != nil {
		return fmt.Errorf("updating namespace %q: %w", ns.Name, err)
	}

	return nil
}

// createNamespace creates the namespace.
func createNamespace(ns Namespace, nsclient corev1typed.NamespaceInterface) error {
	_, err := nsclient.Create(context.TODO(), &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:        ns.Name,
			Labels:      ns.Labels,
			Annotations: ns.Annotations,
		},
	}, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("creating namespace %q: %w", ns.Name, err)
	}

	return nil
}

// ListNamespaces lists the namespaces present in the cluster.
func ListNamespaces(nsclient corev1typed.NamespaceInterface) (*v1.NamespaceList, error) {
	return nsclient.List(context.TODO(), metav1.ListOptions{})
}

// ObjectMetadata uniquely identifies any object in the list of YAML manifests.
type ObjectMetadata struct {
	Version string
	Kind    string
	Name    string
}

// YAMLToUnstructured accepts a Kubernetes manifest in YAML format and returns an object of type
// `unstructured.Unstructured`. This object has many methods that can be used by the consumer to
// extract metadata from the Kubernetes manifest.
func YAMLToUnstructured(yamlObj []byte) (*unstructured.Unstructured, error) {
	u := &unstructured.Unstructured{}

	// Decode YAML into `unstructured.Unstructured`.
	dec := yamlserializer.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	if _, _, err := dec.Decode(yamlObj, nil, u); err != nil {
		return nil, fmt.Errorf("converting config to unstructured.Unstructured: %w", err)
	}

	return u, nil
}

// removeYAMLComments converts YAML to JSON and back again, this removes the comments in the YAML
// and any extra whitespaces spaces.
func removeYAMLComments(yamlObj []byte) ([]byte, error) {
	jsonObj, err := sigyaml.YAMLToJSON(yamlObj)
	if err != nil {
		return nil, fmt.Errorf("converting YAML to JSON: %w", err)
	}

	yamlObj, err = sigyaml.JSONToYAML(jsonObj)
	if err != nil {
		return nil, fmt.Errorf("converting JSON to YAML: %w", err)
	}

	return yamlObj, nil
}

// SplitYAMLDocuments converts a YAML string with multiple YAML docs separated by `---` into unique
// objects and returns those objects as a map.
func SplitYAMLDocuments(yamlObj string) (map[ObjectMetadata]string, error) {
	ret := make(map[ObjectMetadata]string)

	reader := yaml.NewYAMLReader(bufio.NewReader(strings.NewReader(yamlObj)))

	for {
		// Read the YAML document delimited by `---`.
		yamlManifest, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error reading the YAML: %w", err)
		}

		yamlManifest, err = removeYAMLComments(yamlManifest)
		if err != nil {
			return nil, fmt.Errorf("removing YAML comments: %w", err)
		}

		// Check if the YAML is empty.
		if string(yamlManifest) == "null\n" {
			continue
		}

		u, err := YAMLToUnstructured(yamlManifest)
		if err != nil {
			return nil, fmt.Errorf("YAML to unstructured object: %w", err)
		}

		if u.GetAPIVersion() == "" || u.GetKind() == "" {
			return nil, fmt.Errorf("invalid configuration no APIVersion or Kind: %s", string(yamlManifest))
		}

		obj := ObjectMetadata{
			Name:    u.GetName(),
			Kind:    u.GetKind(),
			Version: u.GetAPIVersion(),
		}

		ret[obj] = string(yamlManifest)
	}

	return ret, nil
}
