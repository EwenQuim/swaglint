package main

import (
	"fmt"
	"go/ast"
	"reflect"
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
			pass.Reportf(funcDecl.Pos(), "%s is in code but not in docs", param)
		}
	}
	for _, param := range paramsFromDocs {
		if !slices.Contains(paramsFromCode, param) {
			pass.Reportf(funcDecl.Pos(), "%s is in docs but not in code", param)
		}
	}
}

// checkReturnType returns true if the function has a return type that is not a standard controller
func checkWrongReturnType(funcDecl *ast.FuncDecl) bool {
	lines := funcDecl.Body.List

	for _, line := range lines {
		if assignation, ok := line.(*ast.AssignStmt); ok {
			if len(assignation.Lhs) < 1 {
				continue
			}
			fmt.Printf("lhs: %#v\n", assignation.Lhs[0])
			fmt.Printf("reflect typeof: %v\n", reflect.TypeOf(assignation.Lhs[0]))
			if ident, ok := assignation.Lhs[0].(*ast.Ident); ok {
				fmt.Printf("%#v\n", ident.Obj)
				if ident.Name == "response" {
					fmt.Println(ident.Name, ident.Obj)
					return true
				}
			}
			fmt.Println()
		}
	}
	// lastLine := lines[len(lines)-1]
	// lastExpr, ok := lastLine.(*ast.ExprStmt)
	// if !ok {
	// 	return false
	// }
	// callExpr, ok := lastExpr.X.(*ast.CallExpr)
	// if !ok {
	// 	return false
	// }

	// if len(callExpr.Args) <= 1 {
	// 	return false
	// }

	// argIdent, ok := callExpr.Args[1].(*ast.Ident)
	// if !ok {
	// 	return false
	// }
	// fmt.Printf("=== Working with %#v ===\n", argIdent.Name)
	// fmt.Printf("%#v\n", argIdent)
	// fmt.Printf("%#v\n", argIdent.Obj)
	// fmt.Printf("%#v\n", argIdent.Obj.Decl)

	// rhs, ok := argIdent.Obj.Decl.(*ast.AssignStmt)
	// if !ok {
	// 	return false
	// }

	// fmt.Printf("lhs: %#v\n", rhs.Lhs)
	// for i, lhs := range rhs.Lhs {
	// 	ident, ok := lhs.(*ast.Ident)
	// 	if !ok {
	// 		return false
	// 	}
	// 	fmt.Printf("%#v\n", ident)
	// 	if i == 0 {
	// 		if ident.Obj.Type != nil {
	// 			fmt.Printf("===================\n===================%#v\n", ident.Obj.Type)
	// 		}
	// 	}
	// }
	// fmt.Printf("rhs: %#v\n", rhs.Rhs)
	// for _, rhs := range rhs.Rhs {
	// 	callExpr, ok := rhs.(*ast.CallExpr)
	// 	if !ok {
	// 		return false
	// 	}
	// 	fmt.Printf("%#v\n", callExpr)
	// }
	// fmt.Println()
	// fmt.Println()

	return true
}
