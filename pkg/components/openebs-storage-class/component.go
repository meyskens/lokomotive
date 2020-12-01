// Copyright 2020 The Lokomotive Authors
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

package openebsstorageclass

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"

	"github.com/kinvolk/lokomotive/pkg/components"
	"github.com/kinvolk/lokomotive/pkg/k8sutil"
)

const (
	// Name represents OpenEBS Storage Class component name as it should be referenced in function calls
	// and in configuration.
	Name = "openebs-storage-class"

	poolName = "openebs-storage-pool"
)

type Storageclass struct {
	Name         string   `hcl:"name,label"`
	ReplicaCount int      `hcl:"replica_count,optional"`
	Default      bool     `hcl:"default,optional"`
	Disks        []string `hcl:"disks,optional"`
}
type component struct {
	Storageclasses []Storageclass `hcl:"storage-class,block"`
}

func defaultStorageClass() Storageclass {
	return Storageclass{
		// Name of the storage class
		Name: "openebs-cstor-disk-replica-3",
		// Default replica count value is set to 3
		ReplicaCount: 3,
		// Make the storage class as default
		Default: true,
		// Default disks selection is empty
		Disks: make([]string, 0),
	}
}

// NewConfig returns new OpenEBS Storage Class component configuration with default values set.
//
//nolint:golint
func NewConfig() *component {
	return &component{
		Storageclasses: []Storageclass{},
	}
}

func (c *component) LoadConfig(configBody *hcl.Body, evalContext *hcl.EvalContext) hcl.Diagnostics {
	if configBody == nil {
		return hcl.Diagnostics{}
	}

	if diagnostics := gohcl.DecodeBody(*configBody, evalContext, c); diagnostics != nil {
		return diagnostics
	}
	// if empty config body is provided, default component storage details are still preserved.
	if len(c.Storageclasses) == 0 {
		c.Storageclasses = append(c.Storageclasses, defaultStorageClass())
	}

	if err := c.validateConfig(); err != nil {
		return hcl.Diagnostics{
			&hcl.Diagnostic{
				Severity: hcl.DiagError,
				Summary:  "Validation of the config failed",
				Detail:   fmt.Sprintf("validation failed: %v", err),
			},
		}
	}

	return nil
}

func (c *component) validateConfig() error {
	maxDefaultStorageClass := 0
	for _, sc := range c.Storageclasses {
		if sc.Default == true {
			maxDefaultStorageClass++
		}
		if maxDefaultStorageClass > 1 {
			return fmt.Errorf("cannot have more than one default storage class")
		}
	}

	return nil
}

// TODO: Convert to Helm chart.
func (c *component) RenderManifests() (map[string]string, error) {

	scTmpl, err := template.New(Name).Parse(storageClassTmpl)
	if err != nil {
		return nil, fmt.Errorf("parsing storage class template: %w", err)
	}

	spTmpl, err := template.New(poolName).Parse(storagePoolTmpl)
	if err != nil {
		return nil, fmt.Errorf("parsing storage pool template: %w", err)
	}

	var manifestsMap = make(map[string]string)

	for _, sc := range c.Storageclasses {
		var scBuffer bytes.Buffer
		var spBuffer bytes.Buffer

		if err := scTmpl.Execute(&scBuffer, sc); err != nil {
			return nil, fmt.Errorf("executing storage class %q template: %w", sc.Name, err)
		}

		filename := fmt.Sprintf("%s-%s.yml", Name, sc.Name)
		manifestsMap[filename] = scBuffer.String()

		if err := spTmpl.Execute(&spBuffer, sc); err != nil {
			return nil, fmt.Errorf("executing storage pool %q template: %w", sc.Name, err)
		}

		filename = fmt.Sprintf("%s-%s.yml", poolName, sc.Name)
		manifestsMap[filename] = spBuffer.String()
	}

	return manifestsMap, nil
}

func (c *component) Metadata() components.Metadata {
	return components.Metadata{
		Name: Name,
		// Return the same namespace which the openebs-operator component is using.
		Namespace: k8sutil.Namespace{
			Name: "openebs",
		},
	}
}
