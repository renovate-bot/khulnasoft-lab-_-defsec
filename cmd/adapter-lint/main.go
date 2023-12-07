package main

import (
	"github.com/khulnasoft-lab/defsec/internal/lint/adapter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(adapter.DefaultAnalyzer())
}
