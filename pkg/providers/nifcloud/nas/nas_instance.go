package nas

import (
	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"
)

type NASInstance struct {
	Metadata  defsecTypes.Metadata
	NetworkID defsecTypes.StringValue
}
