package gke

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/google/gke"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckNoLegacyAuthentication(t *testing.T) {
	tests := []struct {
		name     string
		input    gke.GKE
		expected bool
	}{
		{
			name: "Cluster master authentication by certificate",
			input: gke.GKE{
				Clusters: []gke.Cluster{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						MasterAuth: gke.MasterAuth{
							Metadata: defsecTypes.NewTestMetadata(),
							ClientCertificate: gke.ClientCertificate{
								Metadata:         defsecTypes.NewTestMetadata(),
								IssueCertificate: defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
							},
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Cluster master authentication by username/password",
			input: gke.GKE{
				Clusters: []gke.Cluster{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						MasterAuth: gke.MasterAuth{
							Metadata: defsecTypes.NewTestMetadata(),
							ClientCertificate: gke.ClientCertificate{
								Metadata:         defsecTypes.NewTestMetadata(),
								IssueCertificate: defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
							},
							Username: defsecTypes.String("username", defsecTypes.NewTestMetadata()),
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Cluster master authentication by certificate or username/password disabled",
			input: gke.GKE{
				Clusters: []gke.Cluster{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						MasterAuth: gke.MasterAuth{
							Metadata: defsecTypes.NewTestMetadata(),
							ClientCertificate: gke.ClientCertificate{
								Metadata:         defsecTypes.NewTestMetadata(),
								IssueCertificate: defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
							},
							Username: defsecTypes.String("", defsecTypes.NewTestMetadata()),
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
			results := CheckNoLegacyAuthentication.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckNoLegacyAuthentication.Rule().LongID() {
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
