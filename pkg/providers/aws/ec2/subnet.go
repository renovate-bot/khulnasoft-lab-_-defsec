package ec2

import (
	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"
)

type Subnet struct {
	Metadata            defsecTypes.Metadata
	MapPublicIpOnLaunch defsecTypes.BoolValue
}
