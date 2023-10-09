package lambda

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/lambda"
	"github.com/khulnasoft-lab/defsec/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) lambda.Lambda {
	return lambda.Lambda{
		Functions: getFunctions(cfFile),
	}
}
