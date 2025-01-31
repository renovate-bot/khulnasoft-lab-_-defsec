package appservice

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/azure/appservice"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckEnableHttp2(t *testing.T) {
	tests := []struct {
		name     string
		input    appservice.AppService
		expected bool
	}{
		{
			name: "HTTP2 disabled",
			input: appservice.AppService{
				Services: []appservice.Service{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Site: struct {
							EnableHTTP2       defsecTypes.BoolValue
							MinimumTLSVersion defsecTypes.StringValue
						}{
							EnableHTTP2: defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "HTTP2 enabled",
			input: appservice.AppService{
				Services: []appservice.Service{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Site: struct {
							EnableHTTP2       defsecTypes.BoolValue
							MinimumTLSVersion defsecTypes.StringValue
						}{
							EnableHTTP2: defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
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
			testState.Azure.AppService = test.input
			results := CheckEnableHttp2.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckEnableHttp2.Rule().LongID() {
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
