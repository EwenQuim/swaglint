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
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		params := funcDecl.Type.Params.List
		if len(params) != 2 { // [0] must be format (string), [1] must be args (...interface{})
			return true
		}

		// http.ResponseWriter
		firstParamType, ok := params[0].Type.(*ast.SelectorExpr)
		if !ok { // first param type isn't identificator so it can't be of type "string"
			return true
		}
		forstParamLib, ok := firstParamType.X.(*ast.Ident)
		if !ok || forstParamLib.Name != "http" {
			return true
		}
		if firstParamType.Sel.Name != "ResponseWriter" {
			return true
		}

		// *http.Request
		secondParamType, ok := params[1].Type.(*ast.StarExpr)
		if !ok {
			return true
		}
		secondParam, ok := secondParamType.X.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		secondParamdzqid, ok := secondParam.X.(*ast.Ident)
		if !ok || secondParamdzqid.Name != "http" {
			return true
		}
		if secondParam.Sel.Name != "Request" {
			return true
		}

		// no return value
		if funcDecl.Type.Results != nil {
			return true
		}

		// Verify that there are comz
		if funcDecl.Doc == nil {
			pass.Reportf(node.Pos(), "should have a swagger documentation")
			return false
		}

		for _, line := range funcDecl.Doc.List {
			if strings.Contains(line.Text, "@Router") {
				return true
			}
		}
		pass.Reportf(node.Pos(), "no @Router tag found")
		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
