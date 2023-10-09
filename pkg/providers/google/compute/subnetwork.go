package compute

import (
	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"
)

type SubNetwork struct {
	Metadata       defsecTypes.Metadata
	Name           defsecTypes.StringValue
	EnableFlowLogs defsecTypes.BoolValue
}
