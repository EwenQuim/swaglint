package utils

import (
	"go/ast"
	"regexp"
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
	regexp := regexp.MustCompile(`@Param\s+(\w+)\s+` + string(paramType))

	for _, line := range docs.List {

		matches := regexp.FindStringSubmatch(line.Text)
		if len(matches) > 1 {
			paramsFromDocs = append(paramsFromDocs, matches[1])
		}

	}
	return paramsFromDocs
}
