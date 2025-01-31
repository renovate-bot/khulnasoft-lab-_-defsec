package elasticache

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/aws/elasticache"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckEnableAtRestEncryption(t *testing.T) {
	tests := []struct {
		name     string
		input    elasticache.ElastiCache
		expected bool
	}{
		{
			name: "ElastiCache replication group with at-rest encryption disabled",
			input: elasticache.ElastiCache{
				ReplicationGroups: []elasticache.ReplicationGroup{
					{
						Metadata:                defsecTypes.NewTestMetadata(),
						AtRestEncryptionEnabled: defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
					},
				},
			},
			expected: true,
		},
		{
			name: "ElastiCache replication group with at-rest encryption enabled",
			input: elasticache.ElastiCache{
				ReplicationGroups: []elasticache.ReplicationGroup{
					{
						Metadata:                defsecTypes.NewTestMetadata(),
						AtRestEncryptionEnabled: defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
					},
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.AWS.ElastiCache = test.input
			results := CheckEnableAtRestEncryption.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckEnableAtRestEncryption.Rule().LongID() {
					found = true
				}
			}
			if test.expected {
				assert.True(t, found, "Rule should have been found")
			} else {
				assert.False(t, found, "Rule should not have been found")
			}
		})
	}
}
