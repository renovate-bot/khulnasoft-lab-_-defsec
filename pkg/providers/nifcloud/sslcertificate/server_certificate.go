package sslcertificate

import (
	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"
)

type ServerCertificate struct {
	Metadata   defsecTypes.Metadata
	Expiration defsecTypes.TimeValue
}
