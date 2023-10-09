package rds

import (
	"github.com/khulnasoft-lab/defsec/pkg/types"
)

type Classic struct {
	DBSecurityGroups []DBSecurityGroup
}

type DBSecurityGroup struct {
	Metadata types.Metadata
}
