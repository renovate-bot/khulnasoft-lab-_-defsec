package cloudwatch

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/cloudtrail"
	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/cloudwatch"
	"github.com/khulnasoft-lab/defsec/pkg/scan"
	"github.com/khulnasoft-lab/defsec/pkg/state"
	"github.com/stretchr/testify/assert"
)

func TestCheckRouteTableChangeAlarm(t *testing.T) {
	tests := []struct {
		name       string
		cloudtrail cloudtrail.CloudTrail
		cloudwatch cloudwatch.CloudWatch
		expected   bool
	}{
		{
			name: "Multi-region CloudTrail alarms on route table changes",
			cloudtrail: cloudtrail.CloudTrail{
				Trails: []cloudtrail.Trail{
					{
						Metadata:                  defsecTypes.NewTestMetadata(),
						CloudWatchLogsLogGroupArn: defsecTypes.String("arn:aws:cloudwatch:us-east-1:123456789012:log-group:cloudtrail-logging", defsecTypes.NewTestMetadata()),
						IsLogging:                 defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
						IsMultiRegion:             defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
					},
				},
			},
			cloudwatch: cloudwatch.CloudWatch{
				LogGroups: []cloudwatch.LogGroup{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Arn:      defsecTypes.String("arn:aws:cloudwatch:us-east-1:123456789012:log-group:cloudtrail-logging", defsecTypes.NewTestMetadata()),
						MetricFilters: []cloudwatch.MetricFilter{
							{
								Metadata:   defsecTypes.NewTestMetadata(),
								FilterName: defsecTypes.String("RouteTableChange", defsecTypes.NewTestMetadata()),
								FilterPattern: defsecTypes.String(`{($.eventName=CreateRoute) || 
					($.eventName=CreateRouteTable) || ($.eventName=ReplaceRoute) || 
					($.eventName=ReplaceRouteTableAssociation) || ($.eventName=DeleteRouteTable) || 
					($.eventName=DeleteRoute) || ($.eventName=DisassociateRouteTable)}`, defsecTypes.NewTestMetadata()),
							},
						},
					},
				},
				Alarms: []cloudwatch.Alarm{
					{
						Metadata:   defsecTypes.NewTestMetadata(),
						AlarmName:  defsecTypes.String("RouteTableChange", defsecTypes.NewTestMetadata()),
						MetricName: defsecTypes.String("RouteTableChange", defsecTypes.NewTestMetadata()),
						Metrics: []cloudwatch.MetricDataQuery{
							{
								Metadata: defsecTypes.NewTestMetadata(),
								ID:       defsecTypes.String("RouteTableChange", defsecTypes.NewTestMetadata()),
							},
						},
					},
				},
			},
			expected: false,
		},
		{
			name: "Multi-region CloudTrail has no filter for route table changes",
			cloudtrail: cloudtrail.CloudTrail{
				Trails: []cloudtrail.Trail{
					{
						Metadata:                  defsecTypes.NewTestMetadata(),
						CloudWatchLogsLogGroupArn: defsecTypes.String("arn:aws:cloudwatch:us-east-1:123456789012:log-group:cloudtrail-logging", defsecTypes.NewTestMetadata()),
						IsLogging:                 defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
						IsMultiRegion:             defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
					},
				},
			},
			cloudwatch: cloudwatch.CloudWatch{
				LogGroups: []cloudwatch.LogGroup{
					{
						Metadata:      defsecTypes.NewTestMetadata(),
						Arn:           defsecTypes.String("arn:aws:cloudwatch:us-east-1:123456789012:log-group:cloudtrail-logging", defsecTypes.NewTestMetadata()),
						MetricFilters: []cloudwatch.MetricFilter{},
					},
				},
				Alarms: []cloudwatch.Alarm{
					{
						Metadata:  defsecTypes.NewTestMetadata(),
						AlarmName: defsecTypes.String("RouteTableChange", defsecTypes.NewTestMetadata()),
						Metrics: []cloudwatch.MetricDataQuery{
							{},
						},
					},
				},
			},
			expected: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.CloudWatch = test.cloudwatch
			testState.AWS.CloudTrail = test.cloudtrail
			results := requireRouteTableChangeAlarm.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == requireRouteTableChangeAlarm.Rule().LongID() {
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
