package main

import "go/ast"

// isHTTPHandler returns true if the given node is a func declaration
func isHTTPHandler(node ast.Node) (*ast.FuncDecl, bool) {
	funcDecl, ok := node.(*ast.FuncDecl)
	if !ok {
		return nil, false
	}

	params := funcDecl.Type.Params.List
	if len(params) != 2 {
		return nil, false
	}

	// http.ResponseWriter
	firstParamType, ok := params[0].Type.(*ast.SelectorExpr)
	if !ok {
		return nil, false
	}
	forstParamLib, ok := firstParamType.X.(*ast.Ident)
	if !ok || forstParamLib.Name != "http" {
		return nil, false
	}
	if firstParamType.Sel.Name != "ResponseWriter" {
		return nil, false
	}

	// *http.Request
	secondParamType, ok := params[1].Type.(*ast.StarExpr)
	if !ok {
		return nil, false
	}
	secondParam, ok := secondParamType.X.(*ast.SelectorExpr)
	if !ok {
		return nil, false
	}
	secondParamdzqid, ok := secondParam.X.(*ast.Ident)
	if !ok || secondParamdzqid.Name != "http" {
		return nil, false
	}
	if secondParam.Sel.Name != "Request" {
		return nil, false
	}

	// no return value
	if funcDecl.Type.Results != nil {
		return nil, false
	}

	return funcDecl, true
}
