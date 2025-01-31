package compute

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/digitalocean/compute"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckAutoUpgrade(t *testing.T) {
	tests := []struct {
		name     string
		input    compute.Compute
		expected bool
	}{
		{
			name: "Kubernetes cluster auto-upgrade disabled",
			input: compute.Compute{
				KubernetesClusters: []compute.KubernetesCluster{
					{
						Metadata:    defsecTypes.NewTestMetadata(),
						AutoUpgrade: defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
					},
				},
			},
			expected: true,
		},
		{
			name: "Kubernetes cluster auto-upgrade enabled",
			input: compute.Compute{
				KubernetesClusters: []compute.KubernetesCluster{
					{
						Metadata:    defsecTypes.NewTestMetadata(),
						AutoUpgrade: defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
					},
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.DigitalOcean.Compute = test.input
			results := CheckAutoUpgrade.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckAutoUpgrade.Rule().LongID() {
					found = true
				}
			}
			if test.expected {
				assert.True(t, found, "Rule should have been found")
			} else {
				assert.False(t, found, "Rule should not have been found")
			}
		})
	}
}
