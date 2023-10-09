package codebuild

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/codebuild"
	"github.com/khulnasoft-lab/defsec/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) codebuild.CodeBuild {
	return codebuild.CodeBuild{
		Projects: getProjects(cfFile),
	}
}
