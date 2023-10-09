package synapse

import (
	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"
)

type Synapse struct {
	Workspaces []Workspace
}

type Workspace struct {
	Metadata                    defsecTypes.Metadata
	EnableManagedVirtualNetwork defsecTypes.BoolValue
}
