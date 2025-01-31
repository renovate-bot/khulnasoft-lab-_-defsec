package cloudtrail

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/cloudtrail"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckEnableAllRegions(t *testing.T) {
	tests := []struct {
		name     string
		input    cloudtrail.CloudTrail
		expected bool
	}{
		{
			name: "AWS CloudTrail not enabled accross all regions",
			input: cloudtrail.CloudTrail{
				Trails: []cloudtrail.Trail{
					{
						Metadata:      defsecTypes.NewTestMetadata(),
						IsMultiRegion: defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
					},
				},
			},
			expected: true,
		},
		{
			name: "AWS CloudTrail enabled accross all regions",
			input: cloudtrail.CloudTrail{
				Trails: []cloudtrail.Trail{
					{
						Metadata:      defsecTypes.NewTestMetadata(),
						IsMultiRegion: defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
					},
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.CloudTrail = test.input
			results := CheckEnableAllRegions.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckEnableAllRegions.Rule().LongID() {
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
