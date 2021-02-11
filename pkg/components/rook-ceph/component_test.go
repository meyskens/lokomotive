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

package rookceph

import (
	"testing"

	"github.com/hashicorp/hcl/v2"

	"github.com/kinvolk/lokomotive/pkg/components/internal/testutil"
	"github.com/kinvolk/lokomotive/pkg/components/util"
)

func TestEmptyConfig(t *testing.T) {
	c := NewConfig()
	emptyConfig := hcl.EmptyBody()
	evalContext := hcl.EvalContext{}
	diagnostics := c.LoadConfig(&emptyConfig, &evalContext)
	if diagnostics.HasErrors() {
		t.Fatalf("Empty config should not return errors")
	}
}

func TestRenderManifest(t *testing.T) {
	configHCL := `
component "rook-ceph" {
  namespace = "rook-test"

  monitor_count = 3

  node_affinity {
    key      = "node-role.kubernetes.io/storage"
    operator = "Exists"
  }

  node_affinity {
    key      = "storage.lokomotive.io"
    operator = "In"

    values = [
      "foo",
    ]
  }

  toleration {
    key      = "storage.lokomotive.io"
    operator = "Equal"
    value    = "rook-ceph"
    effect   = "NoSchedule"
  }
}
`

	component := NewConfig()

	body, diagnostics := util.GetComponentBody(configHCL, Name)
	if diagnostics != nil {
		t.Fatalf("Error getting component body: %v", diagnostics)
	}

	diagnostics = component.LoadConfig(body, &hcl.EvalContext{})
	if diagnostics.HasErrors() {
		t.Fatalf("Valid config should not return error, got: %s", diagnostics)
	}

	m, err := component.RenderManifests()
	if err != nil {
		t.Fatalf("Rendering manifests with valid config should succeed, got: %s", err)
	}
	if len(m) <= 0 {
		t.Fatalf("Rendered manifests shouldn't be empty")
	}
}

func TestConversion(t *testing.T) {
	testCases := []struct {
		name                 string
		inputConfig          string
		expectedManifestName string
		expected             string
		jsonPath             string
	}{
		{
			name:                 "default reclaim policy",
			inputConfig:          `component "rook-ceph" {}`,
			expectedManifestName: "",
			jsonPath:             "{.reclaimPolicy}",
			expected:             "Retain",
		},
		{
			name: "over-ridden reclaim policy",
			inputConfig: `component "rook-ceph" {
				storage_class {
					reclaim_policy = "Delete"
				}
			}`,
			expectedManifestName: "",
			jsonPath:             "{.reclaimPolicy}",
			expected:             "Delete",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			component := NewConfig()
			m := testutil.RenderManifests(t, component, Name, tc.inputConfig)
			gotConfig := testutil.ConfigFromMap(t, m, tc.expectedManifestName)

			testutil.MatchJSONPathStringValue(t, gotConfig, tc.jsonPath, tc.expected)
		})
	}
}
