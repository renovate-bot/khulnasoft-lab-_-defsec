package compute

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/digitalocean/compute"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckUseSshKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    compute.Compute
		expected bool
	}{
		{
			name: "Droplet missing SSH keys",
			input: compute.Compute{
				Droplets: []compute.Droplet{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						SSHKeys:  []defsecTypes.StringValue{},
					},
				},
			},
			expected: true,
		},
		{
			name: "Droplet with an SSH key provided",
			input: compute.Compute{
				Droplets: []compute.Droplet{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						SSHKeys: []defsecTypes.StringValue{
							defsecTypes.String("my-ssh-key", defsecTypes.NewTestMetadata()),
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
			testState.DigitalOcean.Compute = test.input
			results := CheckUseSshKeys.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckUseSshKeys.Rule().LongID() {
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
