package utils_test

import (
	"go/ast"
	"testing"

	"github.com/EwenQuim/swaglint/analyzer/utils"
	"github.com/stretchr/testify/require"
)

func TestGetParamsFromDoc(t *testing.T) {
	t.Run("no params", func(t *testing.T) {
		params := utils.GetParamsFromDoc(&ast.CommentGroup{
			List: []*ast.Comment{
				{Text: "// @Summary Hello, world!"},
			},
		}, utils.ParamTypeQuery)

		require.Empty(t, params)
	})

	t.Run("no query params", func(t *testing.T) {
		params := utils.GetParamsFromDoc(&ast.CommentGroup{
			List: []*ast.Comment{
				{Text: "// @Param clientID path string true \"Client ID\""},
			},
		}, utils.ParamTypeQuery)

		require.Empty(t, params)
	})

	t.Run("one param", func(t *testing.T) {
		params := utils.GetParamsFromDoc(&ast.CommentGroup{
			List: []*ast.Comment{
				{Text: "// @Param clientID query string true \"Client ID\""},
			},
		}, utils.ParamTypeQuery)

		require.Len(t, params, 1)
		require.Equal(t, "clientID", params[0])
	})

	t.Run("one param with tab", func(t *testing.T) {
		params := utils.GetParamsFromDoc(&ast.CommentGroup{
			List: []*ast.Comment{
				{Text: "// @Param 	clientID query string true \"Client ID\""},
			},
		}, utils.ParamTypeQuery)

		require.Len(t, params, 1)
		require.Equal(t, "clientID", params[0])
	})
}
