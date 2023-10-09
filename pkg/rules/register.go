package rules

import (
	"github.com/khulnasoft-lab/defsec/internal/rules"
	"github.com/khulnasoft-lab/defsec/pkg/framework"
	"github.com/khulnasoft-lab/defsec/pkg/scan"
)

func Register(rule scan.Rule, f scan.CheckFunc) rules.RegisteredRule {
	return rules.Register(rule, f)
}

func GetRegistered(fw ...framework.Framework) (registered []rules.RegisteredRule) {
	return rules.GetFrameworkRules(fw...)
}
