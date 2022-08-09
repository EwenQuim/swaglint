package chianalyzer

import (
	"go/ast"
	"strings"
)

func GetPathParamsFromFunctionBody(body *ast.BlockStmt) []string {
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
				if fun.Sel.Name != "URLParam" {
					continue
				}
				if len(ident.Args) < 2 {
					continue
				}
				arg, ok := ident.Args[1].(*ast.BasicLit)
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
