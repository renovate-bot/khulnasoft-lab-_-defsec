package datafactory

import (
	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"
)

type DataFactory struct {
	DataFactories []Factory
}

type Factory struct {
	Metadata            defsecTypes.Metadata
	EnablePublicNetwork defsecTypes.BoolValue
}
