package parse

import (
	"calculator/frontend/lex"
	"calculator/frontend/parse/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	tcs := []struct {
		Expr string
		AST  ast.Node
	}{
		{"42", ast.Integer(42)},
		{"abc", "abc"},
		{"2+3", ast.BinOp{lex.OpPlus, ast.Integer(2), ast.Integer(3)}},
		{
			"2*3+4", ast.BinOp{
				lex.OpPlus,
				ast.BinOp{lex.OpStar, ast.Integer(2), ast.Integer(3)},
				ast.Integer(4),
			},
		},
		{
			"2+3*4", ast.BinOp{
				lex.OpPlus,
				ast.Integer(2),
				ast.BinOp{lex.OpStar, ast.Integer(3), ast.Integer(4)},
			},
		},
		{
			"(2+3)*4", ast.BinOp{
				lex.OpStar,
				ast.BinOp{lex.OpPlus, ast.Integer(2), ast.Integer(3)},
				ast.Integer(4),
			},
		},
		{
			"-(+5)", ast.UnOp{
				Op: lex.UnMinus,
				Value: ast.UnOp{
					Op:    lex.UnPlus,
					Value: ast.Integer(5),
				},
			},
		},
		{
			"f(x, x+5)", ast.FCall{
				Target: "f",
				Args: []ast.Node{
					"x", ast.BinOp{lex.OpPlus, "x", ast.Integer(5)},
				},
			},
		},
	}

	for _, tc := range tcs {
		testParser(t, tc.Expr, []ast.Node{tc.AST})
	}
}

func testParser(t *testing.T, code string, want ast.Program) {
	lexer := lex.NewLexer(code)
	parser := NewParser(lexer)
	tree, err := parser.Parse()
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, want, tree)
}
