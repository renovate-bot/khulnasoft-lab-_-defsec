package network

import (
	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"
)

type VpnGateway struct {
	Metadata      defsecTypes.Metadata
	SecurityGroup defsecTypes.StringValue
}
