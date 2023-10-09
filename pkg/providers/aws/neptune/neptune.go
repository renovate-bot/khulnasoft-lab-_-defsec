package neptune

import (
	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"
)

type Neptune struct {
	Clusters []Cluster
}

type Cluster struct {
	Metadata         defsecTypes.Metadata
	Logging          Logging
	StorageEncrypted defsecTypes.BoolValue
	KMSKeyID         defsecTypes.StringValue
}

type Logging struct {
	Metadata defsecTypes.Metadata
	Audit    defsecTypes.BoolValue
}
