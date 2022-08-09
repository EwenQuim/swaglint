package analyzer

import (
	"go/ast"
	"strings"

	"github.com/EwenQuim/swaglint/analyzer/stdhttp"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "swaglint",
	Doc:  "Checks that controllers have a swagger documentation",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	inspect := func(node ast.Node) bool {
		funcDecl, ok := stdhttp.IsHTTPHandler(node)
		if !ok {
			return true
		}

		if !checkDocExists(funcDecl) {
			pass.Reportf(node.Pos(), "should have a swagger documentation")
			return false
		}

		missingTags := checkMissingTags(funcDecl)
		if len(missingTags) > 0 {
			pass.Reportf(node.Pos(), "should have the following tags: %s", strings.Join(missingTags, ", "))
			return false
		}

		checkWrongQueryParams(pass, funcDecl)

		checkWrongPathParams(pass, funcDecl)

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
