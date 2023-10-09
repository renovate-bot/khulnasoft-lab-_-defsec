package ecr

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/ecr"
	"github.com/khulnasoft-lab/defsec/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) ecr.ECR {
	return ecr.ECR{
		Repositories: getRepositories(cfFile),
	}
}
