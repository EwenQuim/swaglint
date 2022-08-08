package main

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var Analyzer = &analysis.Analyzer{
	Name: "swaglint",
	Doc:  "Checks that controllers have a swagger documentation",
	Run:  run,
}

func main() {
	singlechecker.Main(Analyzer)
}

func run(pass *analysis.Pass) (any, error) {
	inspect := func(node ast.Node) bool {
		funcDecl, ok := isHTTPHandler(node)
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

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
