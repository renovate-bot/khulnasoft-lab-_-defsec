package github

import (
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/github/branch_protections"
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/github/repositories"
	"github.com/khulnasoft-lab/defsec/internal/adapters/terraform/github/secrets"
	"github.com/khulnasoft-lab/defsec/pkg/providers/github"
	"github.com/khulnasoft-lab/defsec/pkg/terraform"
)

func Adapt(modules terraform.Modules) github.GitHub {
	return github.GitHub{
		Repositories:       repositories.Adapt(modules),
		EnvironmentSecrets: secrets.Adapt(modules),
		BranchProtections:  branch_protections.Adapt(modules),
	}
}
