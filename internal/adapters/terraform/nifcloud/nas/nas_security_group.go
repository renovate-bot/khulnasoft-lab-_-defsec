package nas

import (
	"github.com/khulnasoft-lab/defsec/pkg/providers/nifcloud/nas"
	"github.com/khulnasoft-lab/defsec/pkg/terraform"
	defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"
)

func adaptNASSecurityGroups(modules terraform.Modules) []nas.NASSecurityGroup {
	var nasSecurityGroups []nas.NASSecurityGroup

	for _, resource := range modules.GetResourcesByType("nifcloud_nas_security_group") {
		nasSecurityGroups = append(nasSecurityGroups, adaptNASSecurityGroup(resource))
	}
	return nasSecurityGroups
}

func adaptNASSecurityGroup(resource *terraform.Block) nas.NASSecurityGroup {
	var cidrs []defsecTypes.StringValue

	for _, rule := range resource.GetBlocks("rule") {
		cidrs = append(cidrs, rule.GetAttribute("cidr_ip").AsStringValueOrDefault("", resource))
	}

	return nas.NASSecurityGroup{
		Metadata:    resource.GetMetadata(),
		Description: resource.GetAttribute("description").AsStringValueOrDefault("", resource),
		CIDRs:       cidrs,
	}
}
