package compute

import (
	"github.com/khulnasoft-lab/defsec/pkg/types"
)

type Network struct {
	Metadata    types.Metadata
	Firewall    *Firewall
	Subnetworks []SubNetwork
}
