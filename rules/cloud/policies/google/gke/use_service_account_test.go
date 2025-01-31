package gke

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/google/gke"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckUseServiceAccount(t *testing.T) {
	tests := []struct {
		name     string
		input    gke.GKE
		expected bool
	}{
		{
			name: "Cluster node config with default service account",
			input: gke.GKE{
				Clusters: []gke.Cluster{
					{
						Metadata:              defsecTypes.NewTestMetadata(),
						RemoveDefaultNodePool: defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
						NodeConfig: gke.NodeConfig{
							Metadata:       defsecTypes.NewTestMetadata(),
							ServiceAccount: defsecTypes.String("", defsecTypes.NewTestMetadata()),
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Cluster node config with service account provided",
			input: gke.GKE{
				Clusters: []gke.Cluster{
					{
						Metadata:              defsecTypes.NewTestMetadata(),
						RemoveDefaultNodePool: defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
						NodeConfig: gke.NodeConfig{
							Metadata:       defsecTypes.NewTestMetadata(),
							ServiceAccount: defsecTypes.String("service-account", defsecTypes.NewTestMetadata()),
						},
					},
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.Google.GKE = test.input
			results := CheckUseServiceAccount.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckUseServiceAccount.Rule().LongID() {
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
