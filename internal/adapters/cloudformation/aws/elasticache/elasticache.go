package elasticache

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/elasticache"
	"github.com/khulnasoft-lab/defsec/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) elasticache.ElastiCache {
	return elasticache.ElastiCache{
		Clusters:          getClusterGroups(cfFile),
		ReplicationGroups: getReplicationGroups(cfFile),
		SecurityGroups:    getSecurityGroups(cfFile),
	}
}
