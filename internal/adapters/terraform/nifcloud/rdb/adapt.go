package rdb

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/nifcloud/rdb"
	"github.com/khulnasoft-lab/defsec/pkg/terraform"
)

func Adapt(modules terraform.Modules) rdb.RDB {
	return rdb.RDB{
		DBSecurityGroups: adaptDBSecurityGroups(modules),
		DBInstances:      adaptDBInstances(modules),
	}
}
