package digitalocean

import (
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/digitalocean/compute"
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/digitalocean/spaces"
	"github.com/khulnasoft-lab/defsec/pkg/providers/digitalocean"
	"github.com/khulnasoft-lab/defsec/pkg/terraform"
)

func Adapt(modules terraform.Modules) digitalocean.DigitalOcean {
	return digitalocean.DigitalOcean{
		Compute: compute.Adapt(modules),
		Spaces:  spaces.Adapt(modules),
	}
}
