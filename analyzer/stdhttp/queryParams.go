package stdhttp

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
