package utils

import (
	"go/ast"
	"strings"
)

func GetParamsFromFunctionBody(body *ast.BlockStmt) []string {
	paramsFromCode := make([]string, 0, len(body.List))

	for _, line := range body.List {
		if assignation, ok := line.(*ast.AssignStmt); ok {
			if len(assignation.Lhs) < 1 {
				continue
			}
			if len(assignation.Rhs) < 1 {
				continue
			}

			if ident, ok := assignation.Rhs[0].(*ast.CallExpr); ok {
				fun, ok := ident.Fun.(*ast.SelectorExpr)
				if !ok {
					continue
				}
				if fun.Sel.Name != "FormValue" {
					continue
				}
				if len(ident.Args) < 1 {
					continue
				}
				arg, ok := ident.Args[0].(*ast.BasicLit)
				if !ok {
					continue
				}

				param := strings.TrimFunc(arg.Value, func(r rune) bool {
					return r == '"' || r == '\''
				})
				paramsFromCode = append(paramsFromCode, param)
			}
		}
	}
	return paramsFromCode
}

func GetParamsFromDoc(docs *ast.CommentGroup) []string {
	paramsFromDocs := make([]string, 0, len(docs.List))

	for _, line := range docs.List {
		if !strings.Contains(line.Text, "@Param") {
			continue
		}
		if !strings.Contains(line.Text, "query") {
			continue
		}
		words := removeEmpty(strings.Split(line.Text, " "))
		if len(words) < 3 {
			continue
		}
		paramsFromDocs = append(paramsFromDocs, words[2])

	}
	return paramsFromDocs
}

func removeEmpty(s []string) []string {
	var r []string
	for _, v := range s {
		if v != "" {
			r = append(r, v)
		}
	}
	return r
}
