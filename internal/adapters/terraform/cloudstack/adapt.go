package cloudstack

import (
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/cloudstack/compute"
	"github.com/khulnasoft-lab/defsec/pkg/providers/cloudstack"
	"github.com/khulnasoft-lab/defsec/pkg/terraform"
)

func Adapt(modules terraform.Modules) cloudstack.CloudStack {
	return cloudstack.CloudStack{
		Compute: compute.Adapt(modules),
	}
}
