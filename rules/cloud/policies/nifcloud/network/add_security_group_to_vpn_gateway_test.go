package network

import (
	"testing"

	"github.com/khulnasoft-lab/defsec/pkg/providers/nifcloud/network"
	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckAddSecurityGroupToVpnGateway(t *testing.T) {
	tests := []struct {
		name     string
		input    network.Network
		expected bool
	}{
		{
			name: "NIFCLOUD vpnGateway with no security group provided",
			input: network.Network{
				VpnGateways: []network.VpnGateway{
					{
						Metadata:      defsecTypes.NewTestMetadata(),
						SecurityGroup: defsecTypes.String("", defsecTypes.NewTestMetadata()),
					},
				},
			},
			expected: true,
		},
		{
			name: "NIFCLOUD vpnGateway with security group",
			input: network.Network{
				VpnGateways: []network.VpnGateway{
					{
						Metadata:      defsecTypes.NewTestMetadata(),
						SecurityGroup: defsecTypes.String("some security group", defsecTypes.NewTestMetadata()),
					},
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.Nifcloud.Network = test.input
			results := CheckAddSecurityGroupToVpnGateway.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckAddSecurityGroupToVpnGateway.Rule().LongID() {
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
