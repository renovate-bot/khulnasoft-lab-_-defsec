package tftestutil

import (
	"context"
	"testing"

	"github.com/khulnasoft-lab/defsec/pkg/scanners/terraform/parser"
	"github.com/khulnasoft-lab/defsec/pkg/terraform"

	"github.com/khulnasoft-lab/defsec/test/testutil"
)

func CreateModulesFromSource(t *testing.T, source string, ext string) terraform.Modules {
	fs := testutil.CreateFS(t, map[string]string{
		"source" + ext: source,
	})
	p := parser.New(fs, "", parser.OptionStopOnHCLError(true))
	if err := p.ParseFS(context.TODO(), "."); err != nil {
		t.Fatal(err)
	}
	modules, _, err := p.EvaluateAll(context.TODO())
	if err != nil {
		t.Fatalf("parse error: %s", err)
	}
	return modules
}
