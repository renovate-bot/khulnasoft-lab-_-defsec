package terraform

import (
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/aws"
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/azure"
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/cloudstack"
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/digitalocean"
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/github"
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/google"
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/kubernetes"
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/nifcloud"
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/openstack"
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/oracle"
	"github.com/khulnasoft-lab/defsec/pkg/state"
	"github.com/khulnasoft-lab/defsec/pkg/terraform"
)

func Adapt(modules terraform.Modules) *state.State {
	return &state.State{
		AWS:          aws.Adapt(modules),
		Azure:        azure.Adapt(modules),
		CloudStack:   cloudstack.Adapt(modules),
		DigitalOcean: digitalocean.Adapt(modules),
		GitHub:       github.Adapt(modules),
		Google:       google.Adapt(modules),
		Kubernetes:   kubernetes.Adapt(modules),
		Nifcloud:     nifcloud.Adapt(modules),
		OpenStack:    openstack.Adapt(modules),
		Oracle:       oracle.Adapt(modules),
	}
}
