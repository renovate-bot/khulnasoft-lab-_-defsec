package spaces

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/digitalocean/spaces"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckVersioningEnabled(t *testing.T) {
	tests := []struct {
		name     string
		input    spaces.Spaces
		expected bool
	}{
		{
			name: "Space bucket versioning disabled",
			input: spaces.Spaces{
				Buckets: []spaces.Bucket{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Versioning: spaces.Versioning{
							Metadata: defsecTypes.NewTestMetadata(),
							Enabled:  defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Space bucket versioning enabled",
			input: spaces.Spaces{
				Buckets: []spaces.Bucket{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Versioning: spaces.Versioning{
							Metadata: defsecTypes.NewTestMetadata(),
							Enabled:  defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
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
			testState.DigitalOcean.Spaces = test.input
			results := CheckVersioningEnabled.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckVersioningEnabled.Rule().LongID() {
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
