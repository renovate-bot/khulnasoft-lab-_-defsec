package sql

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/google/sql"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckNoPublicAccess(t *testing.T) {
	tests := []struct {
		name     string
		input    sql.SQL
		expected bool
	}{
		{
			name: "Instance settings set with IPv4 enabled",
			input: sql.SQL{
				Instances: []sql.DatabaseInstance{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Settings: sql.Settings{
							Metadata: defsecTypes.NewTestMetadata(),
							IPConfiguration: sql.IPConfiguration{
								Metadata:   defsecTypes.NewTestMetadata(),
								EnableIPv4: defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
							},
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Instance settings set with IPv4 disabled but public CIDR in authorized networks",
			input: sql.SQL{
				Instances: []sql.DatabaseInstance{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Settings: sql.Settings{
							Metadata: defsecTypes.NewTestMetadata(),
							IPConfiguration: sql.IPConfiguration{
								Metadata:   defsecTypes.NewTestMetadata(),
								EnableIPv4: defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
								AuthorizedNetworks: []struct {
									Name defsecTypes.StringValue
									CIDR defsecTypes.StringValue
								}{
									{
										CIDR: defsecTypes.String("0.0.0.0/0", defsecTypes.NewTestMetadata()),
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
			name: "Instance settings set with IPv4 disabled and private CIDR",
			input: sql.SQL{
				Instances: []sql.DatabaseInstance{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Settings: sql.Settings{
							Metadata: defsecTypes.NewTestMetadata(),
							IPConfiguration: sql.IPConfiguration{
								Metadata:   defsecTypes.NewTestMetadata(),
								EnableIPv4: defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
								AuthorizedNetworks: []struct {
									Name defsecTypes.StringValue
									CIDR defsecTypes.StringValue
								}{
									{
										CIDR: defsecTypes.String("10.0.0.1/24", defsecTypes.NewTestMetadata()),
									},
								},
							},
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
			testState.Google.SQL = test.input
			results := CheckNoPublicAccess.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckNoPublicAccess.Rule().LongID() {
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
