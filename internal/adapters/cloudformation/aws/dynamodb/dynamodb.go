package dynamodb

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/dynamodb"
	"github.com/khulnasoft-lab/defsec/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) dynamodb.DynamoDB {
	return dynamodb.DynamoDB{
		DAXClusters: getClusters(cfFile),
	}
}
