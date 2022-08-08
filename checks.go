package main

import (
	"go/ast"
	"strings"

	"github.com/EwenQuim/swaglint/utils"
	"golang.org/x/exp/slices"
	"golang.org/x/tools/go/analysis"
)

func checkDocExists(funcDecl *ast.FuncDecl) bool {
	return funcDecl.Doc != nil
}

func checkMissingTags(funcDecl *ast.FuncDecl) []string {
	tags := map[string]bool{
		"@Router":  false,
		"@Summary": false,
		"@Tags":    false,
	}

	for _, line := range funcDecl.Doc.List {
		for field := range tags {
			if strings.Contains(line.Text, field) {
				tags[field] = true
			}
		}
	}

	missing := make([]string, 0, len(tags))
	for k, v := range tags {
		if !v {
			missing = append(missing, k)
		}
	}
	return missing
}

func checkWrongQueryParams(pass *analysis.Pass, funcDecl *ast.FuncDecl) {
	paramsFromCode := utils.GetParamsFromFunctionBody(funcDecl.Body)
	paramsFromDocs := utils.GetParamsFromDoc(funcDecl.Doc)

	for _, param := range paramsFromCode {
		if !slices.Contains(paramsFromDocs, param) {
			pass.Reportf(funcDecl.Pos(), "%s is in code but not in docs\n", param)
		}
	}
	for _, param := range paramsFromDocs {
		if !slices.Contains(paramsFromCode, param) {
			pass.Reportf(funcDecl.Pos(), "%s is in docs but not in code\n", param)
		}
	}
}
