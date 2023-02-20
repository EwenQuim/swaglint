package analyzer

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	chianalyzer "github.com/EwenQuim/swaglint/analyzer/chi"
	"github.com/EwenQuim/swaglint/analyzer/stdhttp"
	"github.com/EwenQuim/swaglint/analyzer/utils"
	"golang.org/x/exp/slices"
	"golang.org/x/tools/go/analysis"
)

func checkDocExists(funcDecl *ast.FuncDecl) bool {
	return funcDecl.Doc != nil
}

func checkMissingTags(funcDecl *ast.FuncDecl) []string {
	type tag struct {
		name    string
		present bool
	}

	tags := []tag{
		{name: "@Router", present: false},
		{name: "@Summary", present: false},
		{name: "@Tags", present: false},
	}

	var commentBlock string
	for _, line := range funcDecl.Doc.List {
		commentBlock += line.Text
	}

	for index, tag := range tags {
		if strings.Contains(commentBlock, tag.name) {
			tags[index].present = true
		}
	}

	missing := make([]string, 0, len(tags))
	for _, tag := range tags {
		if !tag.present {
			missing = append(missing, tag.name)
		}
	}
	return missing
}

func checkWrongQueryParams(pass *analysis.Pass, funcDecl *ast.FuncDecl) {
	paramsFromCode := stdhttp.GetParamsFromFunctionBody(funcDecl.Body)
	paramsFromDocs := utils.GetParamsFromDoc(funcDecl.Doc, utils.ParamTypeQuery)

	for _, param := range paramsFromCode {
		if !slices.Contains(paramsFromDocs, param) {
			pass.Reportf(funcDecl.Pos(), "'%s' query param is in code but not in docs", param)
		}
	}
	for _, param := range paramsFromDocs {
		if !slices.Contains(paramsFromCode, param) {
			pass.Reportf(funcDecl.Pos(), "'%s' query param is in docs but not in code", param)
		}
	}
}

func checkWrongPathParams(pass *analysis.Pass, funcDecl *ast.FuncDecl) {
	paramsFromCode := chianalyzer.GetPathParamsFromFunctionBody(funcDecl.Body)
	paramsFromDocs := utils.GetParamsFromDoc(funcDecl.Doc, utils.ParamTypePath)

	for _, param := range paramsFromCode {
		if !slices.Contains(paramsFromDocs, param) {
			pass.Reportf(funcDecl.Pos(), "'%s' path param is in code but not in docs", param)
		}
	}
	for _, param := range paramsFromDocs {
		if !slices.Contains(paramsFromCode, param) {
			pass.Reportf(funcDecl.Pos(), "'%s' path param is in docs but not in code", param)
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
