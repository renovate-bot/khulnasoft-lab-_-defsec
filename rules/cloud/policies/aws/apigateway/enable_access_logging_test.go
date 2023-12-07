package apigateway

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	v1 "github.com/khulnasoft-lab/defsec/pkg/providers/aws/apigateway/v1"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckEnableAccessLogging(t *testing.T) {
	tests := []struct {
		name     string
		input    v1.APIGateway
		expected bool
	}{
		{
			name: "API Gateway stage with no log group ARN",
			input: v1.APIGateway{
				APIs: []v1.API{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Stages: []v1.Stage{
							{
								Metadata: defsecTypes.NewTestMetadata(),
								AccessLogging: v1.AccessLogging{
									Metadata:              defsecTypes.NewTestMetadata(),
									CloudwatchLogGroupARN: defsecTypes.String("", defsecTypes.NewTestMetadata()),
								},
							},
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "API Gateway stage with log group ARN",
			input: v1.APIGateway{
				APIs: []v1.API{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Stages: []v1.Stage{
							{
								Metadata: defsecTypes.NewTestMetadata(),
								AccessLogging: v1.AccessLogging{
									Metadata:              defsecTypes.NewTestMetadata(),
									CloudwatchLogGroupARN: defsecTypes.String("log-group-arn", defsecTypes.NewTestMetadata()),
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
			testState.AWS.APIGateway.V1 = test.input
			results := CheckEnableAccessLogging.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckEnableAccessLogging.Rule().LongID() {
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
