package compute

import (
	"github.com/khulnasoft-lab/defsec/internal/rules"
	"github.com/khulnasoft-lab/defsec/pkg/providers"
	"github.com/khulnasoft-lab/defsec/pkg/scan"
	"github.com/khulnasoft-lab/defsec/pkg/severity"
	"github.com/khulnasoft-lab/defsec/pkg/state"
)

var CheckSecurityGroupHasDescription = rules.Register(
	scan.Rule{
		AVDID:       "AVD-OPNSTK-0005",
		Provider:    providers.OpenStackProvider,
		Service:     "networking",
		ShortCode:   "describe-security-group",
		Summary:     "Missing description for security group.",
		Impact:      "Auditing capability and awareness limited.",
		Resolution:  "Add descriptions for all security groups",
		Explanation: `Security groups should include a description for auditing purposes. Simplifies auditing, debugging, and managing security groups.`,
		Links:       []string{},
		Terraform: &scan.EngineMetadata{
			GoodExamples:        terraformSecurityGroupHasDescriptionGoodExamples,
			BadExamples:         terraformSecurityGroupHasDescriptionBadExamples,
			Links:               terraformSecurityGroupHasDescriptionLinks,
			RemediationMarkdown: terraformSecurityGroupHasDescriptionRemediationMarkdown,
		},
		Severity: severity.Medium,
	},
	func(s *state.State) (results scan.Results) {
		for _, group := range s.OpenStack.Networking.SecurityGroups {
			if group.Metadata.IsUnmanaged() {
				continue
			}
			if group.Description.IsEmpty() {
				results.Add(
					"Security group rule allows egress to multiple public addresses.",
					group.Description,
				)
			} else {
				results.AddPassed(group)
			}
		}
		return
	},
)
