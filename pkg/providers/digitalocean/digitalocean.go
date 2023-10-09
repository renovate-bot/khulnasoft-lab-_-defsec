package digitalocean

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/digitalocean/compute"
	"github.com/khulnasoft-lab/defsec/pkg/providers/digitalocean/spaces"
)

type DigitalOcean struct {
	Compute compute.Compute
	Spaces  spaces.Spaces
}
