package rds

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/rds"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckEnablePerformanceInsightsEncryption(t *testing.T) {
	tests := []struct {
		name     string
		input    rds.RDS
		expected bool
	}{
		{
			name: "RDS Instance with performance insights disabled",
			input: rds.RDS{
				Instances: []rds.Instance{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						PerformanceInsights: rds.PerformanceInsights{
							Metadata: defsecTypes.NewTestMetadata(),
							Enabled:  defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
							KMSKeyID: defsecTypes.String("some-kms-key", defsecTypes.NewTestMetadata()),
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "RDS Instance with performance insights enabled but missing KMS key",
			input: rds.RDS{
				Clusters: []rds.Cluster{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Instances: []rds.ClusterInstance{
							{
								Instance: rds.Instance{
									Metadata: defsecTypes.NewTestMetadata(),
									PerformanceInsights: rds.PerformanceInsights{
										Metadata: defsecTypes.NewTestMetadata(),
										Enabled:  defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
										KMSKeyID: defsecTypes.String("", defsecTypes.NewTestMetadata()),
									},
								},
							},
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "RDS Instance with performance insights enabled and KMS key provided",
			input: rds.RDS{
				Instances: []rds.Instance{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						PerformanceInsights: rds.PerformanceInsights{
							Metadata: defsecTypes.NewTestMetadata(),
							Enabled:  defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
							KMSKeyID: defsecTypes.String("some-kms-key", defsecTypes.NewTestMetadata()),
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
			testState.AWS.RDS = test.input
			results := CheckEnablePerformanceInsightsEncryption.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() != scan.StatusPassed && result.Rule().LongID() == CheckEnablePerformanceInsightsEncryption.Rule().LongID() {
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
