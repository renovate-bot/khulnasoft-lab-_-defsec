package cloudformation

import (
	"github.com/khulnasoft-lab/defsec/internal/adapters/cloudformation/aws"
	"github.com/khulnasoft-lab/defsec/pkg/scanners/cloudformation/parser"
	"github.com/khulnasoft-lab/defsec/pkg/state"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) *state.State {
	return &state.State{
		AWS: aws.Adapt(cfFile),
	}
}
