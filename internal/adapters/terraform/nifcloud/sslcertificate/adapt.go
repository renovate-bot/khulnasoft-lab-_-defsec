package sslcertificate

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/nifcloud/sslcertificate"
	"github.com/khulnasoft-lab/defsec/pkg/terraform"
)

func Adapt(modules terraform.Modules) sslcertificate.SSLCertificate {
	return sslcertificate.SSLCertificate{
		ServerCertificates: adaptServerCertificates(modules),
	}
}
