package neptune

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/neptune"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckEnableLogExport(t *testing.T) {
	tests := []struct {
		name     string
		input    neptune.Neptune
		expected bool
	}{
		{
			name: "Neptune Cluster with audit logging disabled",
			input: neptune.Neptune{
				Clusters: []neptune.Cluster{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Logging: neptune.Logging{
							Metadata: defsecTypes.NewTestMetadata(),
							Audit:    defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Neptune Cluster with audit logging enabled",
			input: neptune.Neptune{
				Clusters: []neptune.Cluster{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Logging: neptune.Logging{
							Metadata: defsecTypes.NewTestMetadata(),
							Audit:    defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
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
			testState.AWS.Neptune = test.input
			results := CheckEnableLogExport.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckEnableLogExport.Rule().LongID() {
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
