package s3

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/s3"
	"github.com/khulnasoft-lab/defsec/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) s3.S3 {
	return s3.S3{
		Buckets: getBuckets(cfFile),
	}
}
