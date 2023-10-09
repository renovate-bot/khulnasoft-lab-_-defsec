package ssm

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/ssm"
	"github.com/khulnasoft-lab/defsec/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) ssm.SSM {
	return ssm.SSM{
		Secrets: getSecrets(cfFile),
	}
}
