package main

import (
	"go/ast"
	"strings"
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
