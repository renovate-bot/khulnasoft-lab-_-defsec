package cloudfront

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/cloudfront"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckUseSecureTlsPolicy(t *testing.T) {
	tests := []struct {
		name     string
		input    cloudfront.Cloudfront
		expected bool
	}{
		{
			name: "CloudFront distribution using TLS v1.0",
			input: cloudfront.Cloudfront{
				Distributions: []cloudfront.Distribution{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						ViewerCertificate: cloudfront.ViewerCertificate{
							Metadata:               defsecTypes.NewTestMetadata(),
							MinimumProtocolVersion: defsecTypes.String("TLSv1.0", defsecTypes.NewTestMetadata()),
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "CloudFront distribution using TLS v1.2",
			input: cloudfront.Cloudfront{
				Distributions: []cloudfront.Distribution{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						ViewerCertificate: cloudfront.ViewerCertificate{
							Metadata:               defsecTypes.NewTestMetadata(),
							MinimumProtocolVersion: defsecTypes.String(cloudfront.ProtocolVersionTLS1_2, defsecTypes.NewTestMetadata()),
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
			testState.AWS.Cloudfront = test.input
			results := CheckUseSecureTlsPolicy.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckUseSecureTlsPolicy.Rule().LongID() {
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
