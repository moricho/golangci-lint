package golinters

import (
	"github.com/moricho/tparallel"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
)

func NewTparallel() *goanalysis.Linter {
	analyzers := []*analysis.Analyzer{
		tparallel.Analyzer,
	}

	return goanalysis.NewLinter(
		"tparallel",
		"checks whether the usage of t.Parallel() method is appropriate in your test codes",
		analyzers,
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
