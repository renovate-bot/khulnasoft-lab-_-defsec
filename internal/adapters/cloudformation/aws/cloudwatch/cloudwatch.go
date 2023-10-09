package cloudwatch

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/cloudwatch"
	"github.com/khulnasoft-lab/defsec/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) cloudwatch.CloudWatch {
	return cloudwatch.CloudWatch{
		LogGroups: getLogGroups(cfFile),
		Alarms:    nil,
	}
}
