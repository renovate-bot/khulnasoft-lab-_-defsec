package redshift

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/redshift"
	"github.com/khulnasoft-lab/defsec/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) redshift.Redshift {
	return redshift.Redshift{
		Clusters:          getClusters(cfFile),
		SecurityGroups:    getSecurityGroups(cfFile),
		ClusterParameters: getParameters(cfFile),
		ReservedNodes:     nil,
	}
}
