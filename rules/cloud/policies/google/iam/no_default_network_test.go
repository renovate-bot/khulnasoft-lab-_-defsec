package iam

import (
	"testing"

	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

	"github.com/khulnasoft-lab/defsec/pkg/state"

	"github.com/khulnasoft-lab/defsec/pkg/providers/google/iam"
	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/stretchr/testify/assert"
)

func TestCheckNoDefaultNetwork(t *testing.T) {
	tests := []struct {
		name     string
		input    iam.IAM
		expected bool
	}{
		{
			name: "Project automatic network creation enabled",
			input: iam.IAM{
				Organizations: []iam.Organization{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Projects: []iam.Project{
							{
								Metadata:          defsecTypes.NewTestMetadata(),
								AutoCreateNetwork: defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
							},
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Project automatic network creation enabled #2",
			input: iam.IAM{
				Organizations: []iam.Organization{
					{
						Metadata: defsecTypes.NewTestMetadata(),

						Folders: []iam.Folder{
							{
								Metadata: defsecTypes.NewTestMetadata(),
								Projects: []iam.Project{
									{
										Metadata:          defsecTypes.NewTestMetadata(),
										AutoCreateNetwork: defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
									},
								},
								Folders: []iam.Folder{
									{
										Metadata: defsecTypes.NewTestMetadata(),
										Projects: []iam.Project{
											{
												Metadata:          defsecTypes.NewTestMetadata(),
												AutoCreateNetwork: defsecTypes.Bool(true, defsecTypes.NewTestMetadata()),
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: true,
		},
		{
			name: "Project automatic network creation disabled",
			input: iam.IAM{
				Organizations: []iam.Organization{
					{
						Metadata: defsecTypes.NewTestMetadata(),
						Projects: []iam.Project{
							{
								Metadata:          defsecTypes.NewTestMetadata(),
								AutoCreateNetwork: defsecTypes.Bool(false, defsecTypes.NewTestMetadata()),
							},
						},
					},
				},
			},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var testState state.State
			testState.Google.IAM = test.input
			results := CheckNoDefaultNetwork.Evaluate(&testState)
			var found bool
			for _, result := range results {
				if result.Status() == scan.StatusFailed && result.Rule().LongID() == CheckNoDefaultNetwork.Rule().LongID() {
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
