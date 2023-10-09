package sqs

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/sqs"
	"github.com/khulnasoft-lab/defsec/pkg/scanners/cloudformation/parser"
)

// Adapt ...
func Adapt(cfFile parser.FileContext) sqs.SQS {
	return sqs.SQS{
		Queues: getQueues(cfFile),
	}
}
