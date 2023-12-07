package arm

import (
	"context"

	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/appservice"
	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/authorization"
	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/compute"
	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/container"
	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/database"
	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/datafactory"
	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/datalake"
	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/keyvault"
	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/monitor"
	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/network"
	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/securitycenter"
	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/storage"
	"github.com/khulnasoft-lab/defsec/internal/adapters/arm/synapse"

	"github.com/khulnasoft-lab/defsec/pkg/providers/azure"
	scanner "github.com/khulnasoft-lab/defsec/pkg/scanners/azure"
	"github.com/khulnasoft-lab/defsec/pkg/state"
)

// Adapt ...
func Adapt(ctx context.Context, deployment scanner.Deployment) *state.State {
	return &state.State{
		Azure: adaptAzure(deployment),
	}
}

func adaptAzure(deployment scanner.Deployment) azure.Azure {

	return azure.Azure{
		AppService:     appservice.Adapt(deployment),
		Authorization:  authorization.Adapt(deployment),
		Compute:        compute.Adapt(deployment),
		Container:      container.Adapt(deployment),
		Database:       database.Adapt(deployment),
		DataFactory:    datafactory.Adapt(deployment),
		DataLake:       datalake.Adapt(deployment),
		KeyVault:       keyvault.Adapt(deployment),
		Monitor:        monitor.Adapt(deployment),
		Network:        network.Adapt(deployment),
		SecurityCenter: securitycenter.Adapt(deployment),
		Storage:        storage.Adapt(deployment),
		Synapse:        synapse.Adapt(deployment),
	}

}
