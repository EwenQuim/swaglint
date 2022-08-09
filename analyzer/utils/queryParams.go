package utils

import (
	"go/ast"
	"strings"
)

type ParamType string

const (
	ParamTypeQuery  ParamType = "query"
	ParamTypePath   ParamType = "path"
	ParamTypeHeader ParamType = "header"
	ParamTypeBody   ParamType = "body"
)

func GetParamsFromDoc(docs *ast.CommentGroup, paramType ParamType) []string {
	paramsFromDocs := make([]string, 0, len(docs.List))

	for _, line := range docs.List {
		if !strings.Contains(line.Text, "@Param") {
			continue
		}
		if !strings.Contains(line.Text, string(paramType)) {
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
