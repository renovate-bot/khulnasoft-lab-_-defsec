package test

import (
	"fmt"
	"testing"

	"github.com/khulnasoft-lab/defsec/pkg/terraform"

	"github.com/khulnasoft-lab/defsec/pkg/severity"

	"github.com/khulnasoft-lab/defsec/pkg/scan"

	"github.com/khulnasoft-lab/defsec/internal/rules"

	"github.com/khulnasoft-lab/defsec/pkg/providers"

	"github.com/stretchr/testify/assert"
)

var exampleRule = scan.Rule{
	Provider:  providers.AWSProvider,
	Service:   "service",
	ShortCode: "abc123",
	Aliases:   []string{"aws-other-abc123"},
	Severity:  severity.High,
	CustomChecks: scan.CustomChecks{
		Terraform: &scan.TerraformCustomCheck{
			RequiredLabels: []string{"bad"},
			Check: func(resourceBlock *terraform.Block, _ *terraform.Module) (results scan.Results) {
				attr := resourceBlock.GetAttribute("secure")
				if attr.IsNil() {
					results.Add("example problem", resourceBlock)
				}
				if attr.IsFalse() {
					results.Add("example problem", attr)
				}
				return
			},
		},
	},
}

func Test_IgnoreAll(t *testing.T) {

	var testCases = []struct {
		name         string
		inputOptions string
		assertLength int
	}{
		{name: "IgnoreAll", inputOptions: `
resource "bad" "my-rule" {
    secure = false // tfsec:ignore:*
}
`, assertLength: 0},
		{name: "IgnoreLineAboveTheBlock", inputOptions: `
// tfsec:ignore:*
resource "bad" "my-rule" {
   secure = false 
}
`, assertLength: 0},
		{name: "IgnoreLineAboveTheBlockMatchingParamBool", inputOptions: `
// tfsec:ignore:*[secure=false]
resource "bad" "my-rule" {
   secure = false
}
`, assertLength: 0},
		{name: "IgnoreLineAboveTheBlockNotMatchingParamBool", inputOptions: `
// tfsec:ignore:*[secure=true]
resource "bad" "my-rule" {
   secure = false 
}
`, assertLength: 1},
		{name: "IgnoreLineAboveTheBlockMatchingParamString", inputOptions: `
// tfsec:ignore:*[name=myrule]
resource "bad" "my-rule" {
    name = "myrule"
    secure = false
}
`, assertLength: 0},
		{name: "IgnoreLineAboveTheBlockNotMatchingParamString", inputOptions: `
// tfsec:ignore:*[name=myrule2]
resource "bad" "my-rule" {
    name = "myrule"
    secure = false 
}
`, assertLength: 1},
		{name: "IgnoreLineAboveTheBlockMatchingParamInt", inputOptions: `
// tfsec:ignore:*[port=123]
resource "bad" "my-rule" {
   secure = false
   port = 123
}
`, assertLength: 0},
		{name: "IgnoreLineAboveTheBlockNotMatchingParamInt", inputOptions: `
// tfsec:ignore:*[port=456]
resource "bad" "my-rule" {
   secure = false 
   port = 123
}
`, assertLength: 1},
		{name: "IgnoreLineStackedAboveTheBlock", inputOptions: `
// tfsec:ignore:*
// tfsec:ignore:a
// tfsec:ignore:b
// tfsec:ignore:c
// tfsec:ignore:d
resource "bad" "my-rule" {
   secure = false 
}
`, assertLength: 0},
		{name: "IgnoreLineStackedAboveTheBlockWithoutMatch", inputOptions: `
#tfsec:ignore:*

#tfsec:ignore:x
#tfsec:ignore:a
#tfsec:ignore:b
#tfsec:ignore:c
#tfsec:ignore:d
resource "bad" "my-rule" {
   secure = false 
}
`, assertLength: 1},
		{name: "IgnoreLineStackedAboveTheBlockWithHashesWithoutSpaces", inputOptions: `
#tfsec:ignore:*
#tfsec:ignore:a
#tfsec:ignore:b
#tfsec:ignore:c
#tfsec:ignore:d
resource "bad" "my-rule" {
   secure = false 
}
`, assertLength: 0},
		{name: "IgnoreLineStackedAboveTheBlockWithoutSpaces", inputOptions: `
//tfsec:ignore:*
//tfsec:ignore:a
//tfsec:ignore:b
//tfsec:ignore:c
//tfsec:ignore:d
resource "bad" "my-rule" {
   secure = false 
}
`, assertLength: 0},
		{name: "IgnoreLineAboveTheLine", inputOptions: `
resource "bad" "my-rule" {
	# tfsec:ignore:aws-service-abc123
    secure = false
}
`, assertLength: 0},
		{name: "IgnoreWithExpDateIfDateBreachedThenDontIgnore", inputOptions: `
resource "bad" "my-rule" {
    secure = false # tfsec:ignore:aws-service-abc123:exp:2000-01-02
}
`, assertLength: 1},
		{name: "IgnoreWithExpDateIfDateNotBreachedThenIgnoreIgnore", inputOptions: `
resource "bad" "my-rule" {
    secure = false # tfsec:ignore:aws-service-abc123:exp:2221-01-02
}
`, assertLength: 0},
		{name: "IgnoreWithExpDateIfDateInvalidThenDropTheIgnore", inputOptions: `
resource "bad" "my-rule" {
   secure = false # tfsec:ignore:aws-service-abc123:exp:2221-13-02
}
`, assertLength: 1},
		{name: "IgnoreAboveResourceBlockWithExpDateIfDateNotBreachedThenIgnoreIgnore", inputOptions: `
#tfsec:ignore:aws-service-abc123:exp:2221-01-02
resource "bad" "my-rule" {
}
`, assertLength: 0},
		{name: "IgnoreAboveResourceBlockWithExpDateAndMultipleIgnoresIfDateNotBreachedThenIgnoreIgnore", inputOptions: `
# tfsec:ignore:aws-service-abc123:exp:2221-01-02
resource "bad" "my-rule" {
	
}
`, assertLength: 0},
		{name: "IgnoreForImpliedIAMResource", inputOptions: `
terraform {
required_version = "~> 1.1.6"

required_providers {
aws = {
source  = "hashicorp/aws"
version = "~> 3.48"
}
}
}

# Retrieve an IAM group defined outside of this Terraform config.

# tfsec:ignore:aws-iam-enforce-mfa
data "aws_iam_group" "externally_defined_group" {
group_name = "group-name" # tfsec:ignore:aws-iam-enforce-mfa
}

# Create an IAM policy and attach it to the group.

# tfsec:ignore:aws-iam-enforce-mfa
resource "aws_iam_policy" "test_policy" {
name   = "test-policy" # tfsec:ignore:aws-iam-enforce-mfa
policy = data.aws_iam_policy_document.test_policy.json # tfsec:ignore:aws-iam-enforce-mfa
}

# tfsec:ignore:aws-iam-enforce-mfa
resource "aws_iam_group_policy_attachment" "test_policy_attachment" {
group      = data.aws_iam_group.externally_defined_group.group_name # tfsec:ignore:aws-iam-enforce-mfa
policy_arn = aws_iam_policy.test_policy.arn # tfsec:ignore:aws-iam-enforce-mfa
}

# tfsec:ignore:aws-iam-enforce-mfa
data "aws_iam_policy_document" "test_policy" {
statement {
sid = "PublishToCloudWatch" # tfsec:ignore:aws-iam-enforce-mfa
actions = [
"cloudwatch:PutMetricData", # tfsec:ignore:aws-iam-enforce-mfa
]
resources = ["*"] # tfsec:ignore:aws-iam-enforce-mfa
}
}
`, assertLength: 0},
		{name: "TunnelIgnoreAll", inputOptions: `
resource "bad" "my-rule" {
    secure = false // tunnel:ignore:*
}
`, assertLength: 0},
		{name: "TunnelIgnoreLineAboveTheBlock", inputOptions: `
// tunnel:ignore:*
resource "bad" "my-rule" {
   secure = false 
}
`, assertLength: 0},
		{name: "TunnelIgnoreLineAboveTheBlockMatchingParamBool", inputOptions: `
// tunnel:ignore:*[secure=false]
resource "bad" "my-rule" {
   secure = false
}
`, assertLength: 0},
		{name: "TunnelIgnoreLineAboveTheBlockNotMatchingParamBool", inputOptions: `
// tunnel:ignore:*[secure=true]
resource "bad" "my-rule" {
   secure = false 
}
`, assertLength: 1},
		{name: "TunnelIgnoreLineAboveTheBlockMatchingParamString", inputOptions: `
// tunnel:ignore:*[name=myrule]
resource "bad" "my-rule" {
    name = "myrule"
    secure = false
}
`, assertLength: 0},
		{name: "TunnelIgnoreLineAboveTheBlockNotMatchingParamString", inputOptions: `
// tunnel:ignore:*[name=myrule2]
resource "bad" "my-rule" {
    name = "myrule"
    secure = false 
}
`, assertLength: 1},
		{name: "TunnelIgnoreLineAboveTheBlockMatchingParamInt", inputOptions: `
// tunnel:ignore:*[port=123]
resource "bad" "my-rule" {
   secure = false
   port = 123
}
`, assertLength: 0},
		{name: "TunnelIgnoreLineAboveTheBlockNotMatchingParamInt", inputOptions: `
// tunnel:ignore:*[port=456]
resource "bad" "my-rule" {
   secure = false 
   port = 123
}
`, assertLength: 1},
		{name: "TunnelIgnoreLineStackedAboveTheBlock", inputOptions: `
// tunnel:ignore:*
// tunnel:ignore:a
// tunnel:ignore:b
// tunnel:ignore:c
// tunnel:ignore:d
resource "bad" "my-rule" {
   secure = false 
}
`, assertLength: 0},
		{name: "TunnelIgnoreLineStackedAboveTheBlockWithoutMatch", inputOptions: `
#tunnel:ignore:*

#tunnel:ignore:x
#tunnel:ignore:a
#tunnel:ignore:b
#tunnel:ignore:c
#tunnel:ignore:d
resource "bad" "my-rule" {
   secure = false 
}
`, assertLength: 1},
		{name: "TunnelIgnoreLineStackedAboveTheBlockWithHashesWithoutSpaces", inputOptions: `
#tunnel:ignore:*
#tunnel:ignore:a
#tunnel:ignore:b
#tunnel:ignore:c
#tunnel:ignore:d
resource "bad" "my-rule" {
   secure = false 
}
`, assertLength: 0},
		{name: "TunnelIgnoreLineStackedAboveTheBlockWithoutSpaces", inputOptions: `
//tunnel:ignore:*
//tunnel:ignore:a
//tunnel:ignore:b
//tunnel:ignore:c
//tunnel:ignore:d
resource "bad" "my-rule" {
   secure = false 
}
`, assertLength: 0},
		{name: "TunnelIgnoreLineAboveTheLine", inputOptions: `
resource "bad" "my-rule" {
	# tunnel:ignore:aws-service-abc123
    secure = false
}
`, assertLength: 0},
		{name: "TunnelIgnoreWithExpDateIfDateBreachedThenDontIgnore", inputOptions: `
resource "bad" "my-rule" {
    secure = false # tunnel:ignore:aws-service-abc123:exp:2000-01-02
}
`, assertLength: 1},
		{name: "TunnelIgnoreWithExpDateIfDateNotBreachedThenIgnoreIgnore", inputOptions: `
resource "bad" "my-rule" {
    secure = false # tunnel:ignore:aws-service-abc123:exp:2221-01-02
}
`, assertLength: 0},
		{name: "TunnelIgnoreWithExpDateIfDateInvalidThenDropTheIgnore", inputOptions: `
resource "bad" "my-rule" {
   secure = false # tunnel:ignore:aws-service-abc123:exp:2221-13-02
}
`, assertLength: 1},
		{name: "TunnelIgnoreAboveResourceBlockWithExpDateIfDateNotBreachedThenIgnoreIgnore", inputOptions: `
#tunnel:ignore:aws-service-abc123:exp:2221-01-02
resource "bad" "my-rule" {
}
`, assertLength: 0},
		{name: "TunnelIgnoreAboveResourceBlockWithExpDateAndMultipleIgnoresIfDateNotBreachedThenIgnoreIgnore", inputOptions: `
# tunnel:ignore:aws-service-abc123:exp:2221-01-02
resource "bad" "my-rule" {
	
}
`, assertLength: 0},
		{name: "TunnelIgnoreForImpliedIAMResource", inputOptions: `
terraform {
required_version = "~> 1.1.6"

required_providers {
aws = {
source  = "hashicorp/aws"
version = "~> 3.48"
}
}
}

# Retrieve an IAM group defined outside of this Terraform config.

# tunnel:ignore:aws-iam-enforce-mfa
data "aws_iam_group" "externally_defined_group" {
group_name = "group-name" # tunnel:ignore:aws-iam-enforce-mfa
}

# Create an IAM policy and attach it to the group.

# tunnel:ignore:aws-iam-enforce-mfa
resource "aws_iam_policy" "test_policy" {
name   = "test-policy" # tunnel:ignore:aws-iam-enforce-mfa
policy = data.aws_iam_policy_document.test_policy.json # tunnel:ignore:aws-iam-enforce-mfa
}

# tunnel:ignore:aws-iam-enforce-mfa
resource "aws_iam_group_policy_attachment" "test_policy_attachment" {
group      = data.aws_iam_group.externally_defined_group.group_name # tunnel:ignore:aws-iam-enforce-mfa
policy_arn = aws_iam_policy.test_policy.arn # tunnel:ignore:aws-iam-enforce-mfa
}

# tunnel:ignore:aws-iam-enforce-mfa
data "aws_iam_policy_document" "test_policy" {
statement {
sid = "PublishToCloudWatch" # tunnel:ignore:aws-iam-enforce-mfa
actions = [
"cloudwatch:PutMetricData", # tunnel:ignore:aws-iam-enforce-mfa
]
resources = ["*"] # tunnel:ignore:aws-iam-enforce-mfa
}
}
`, assertLength: 0}}

	reg := rules.Register(exampleRule, nil)
	defer rules.Deregister(reg)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			results := scanHCL(t, tc.inputOptions)
			assert.Len(t, results.GetFailed(), tc.assertLength)
		})
	}
}

func Test_IgnoreIgnoreWithExpiryAndWorkspaceAndWorkspaceSupplied(t *testing.T) {
	reg := rules.Register(exampleRule, nil)
	defer rules.Deregister(reg)

	results := scanHCLWithWorkspace(t, `
# tfsec:ignore:aws-service-abc123:exp:2221-01-02:ws:testworkspace
resource "bad" "my-rule" {
}
`, "testworkspace")
	assert.Len(t, results.GetFailed(), 0)
}

func Test_IgnoreInline(t *testing.T) {
	reg := rules.Register(exampleRule, nil)
	defer rules.Deregister(reg)

	results := scanHCL(t, fmt.Sprintf(`
	resource "bad" "sample" {
		  secure = false # tfsec:ignore:%s
	}
	  `, exampleRule.LongID()))
	assert.Len(t, results.GetFailed(), 0)
}

func Test_IgnoreIgnoreWithExpiryAndWorkspaceButWrongWorkspaceSupplied(t *testing.T) {
	reg := rules.Register(exampleRule, nil)
	defer rules.Deregister(reg)

	results := scanHCLWithWorkspace(t, `
# tfsec:ignore:aws-service-abc123:exp:2221-01-02:ws:otherworkspace
resource "bad" "my-rule" {
	
}
`, "testworkspace")
	assert.Len(t, results.GetFailed(), 1)
}

func Test_IgnoreWithAliasCodeStillIgnored(t *testing.T) {
	reg := rules.Register(exampleRule, nil)
	defer rules.Deregister(reg)

	results := scanHCLWithWorkspace(t, `
# tfsec:ignore:aws-other-abc123
resource "bad" "my-rule" {
	
}
`, "testworkspace")
	assert.Len(t, results.GetFailed(), 0)
}

func Test_TunnelIgnoreIgnoreWithExpiryAndWorkspaceAndWorkspaceSupplied(t *testing.T) {
	reg := rules.Register(exampleRule, nil)
	defer rules.Deregister(reg)

	results := scanHCLWithWorkspace(t, `
# tunnel:ignore:aws-service-abc123:exp:2221-01-02:ws:testworkspace
resource "bad" "my-rule" {
}
`, "testworkspace")
	assert.Len(t, results.GetFailed(), 0)
}

func Test_TunnelIgnoreIgnoreWithExpiryAndWorkspaceButWrongWorkspaceSupplied(t *testing.T) {
	reg := rules.Register(exampleRule, nil)
	defer rules.Deregister(reg)

	results := scanHCLWithWorkspace(t, `
# tunnel:ignore:aws-service-abc123:exp:2221-01-02:ws:otherworkspace
resource "bad" "my-rule" {
	
}
`, "testworkspace")
	assert.Len(t, results.GetFailed(), 1)
}

func Test_TunnelIgnoreWithAliasCodeStillIgnored(t *testing.T) {
	reg := rules.Register(exampleRule, nil)
	defer rules.Deregister(reg)

	results := scanHCLWithWorkspace(t, `
# tunnel:ignore:aws-other-abc123
resource "bad" "my-rule" {
	
}
`, "testworkspace")
	assert.Len(t, results.GetFailed(), 0)
}

func Test_TunnelIgnoreInline(t *testing.T) {
	reg := rules.Register(exampleRule, nil)
	defer rules.Deregister(reg)

	results := scanHCL(t, fmt.Sprintf(`
	resource "bad" "sample" {
		  secure = false # tunnel:ignore:%s
	}
	  `, exampleRule.LongID()))
	assert.Len(t, results.GetFailed(), 0)
}
