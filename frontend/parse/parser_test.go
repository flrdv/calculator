package parse

import (
	"calculator/frontend/lex"
	"calculator/frontend/parse/ast"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	tcs := []struct {
		Name string
		Expr string
		Want ast.Node
	}{
		{
			Name: "single number",
			Expr: "42",
			Want: ast.Integer(42),
		},
		{
			Name: "single id",
			Expr: "abc",
			Want: "abc",
		},
		{
			Name: "single binary operation",
			Expr: "2+3",
			Want: ast.BinOp{Op: lex.OpPlus, Left: ast.Integer(2), Right: ast.Integer(3)},
		},
		{
			Name: "priority of multiplication over addition",
			Expr: "2*3+4",
			Want: ast.BinOp{
				Op:    lex.OpPlus,
				Left:  ast.BinOp{Op: lex.OpStar, Left: ast.Integer(2), Right: ast.Integer(3)},
				Right: ast.Integer(4),
			},
		},
		{
			Name: "priority of multiplication over addition (but a bit different)",
			Expr: "2+3*4",
			Want: ast.BinOp{
				Op:    lex.OpPlus,
				Left:  ast.Integer(2),
				Right: ast.BinOp{Op: lex.OpStar, Left: ast.Integer(3), Right: ast.Integer(4)},
			},
		},
		{
			Name: "parenthesis priority",
			Expr: "(2+3)*4",
			Want: ast.BinOp{
				Op:    lex.OpStar,
				Left:  ast.BinOp{Op: lex.OpPlus, Left: ast.Integer(2), Right: ast.Integer(3)},
				Right: ast.Integer(4),
			},
		},
		{
			Name: "unary with expression in parenthesis",
			Expr: "-(+5)",
			Want: ast.UnOp{
				Op: lex.UnMinus,
				Value: ast.UnOp{
					Op:    lex.UnPlus,
					Value: ast.Integer(5),
				},
			},
		},
		{
			Name: "pass expression via arguments",
			Expr: "f(x, x+5)",
			Want: ast.FCall{
				Target: "f",
				Args: []ast.Node{
					"x", ast.BinOp{Op: lex.OpPlus, Left: "x", Right: ast.Integer(5)},
				},
			},
		},
		{
			Name: "call the result of another function call",
			Expr: "f(x)(y)",
			Want: ast.FCall{
				Target: ast.FCall{
					Target: "f",
					Args:   []ast.Node{"x"},
				},
				Args: []ast.Node{"y"},
			},
		},
	}

	for _, tc := range tcs {
		testParser(t, fmtTestName(tc.Name, tc.Expr), tc.Expr, []ast.Node{tc.Want})
	}
}

func testParser(t *testing.T, name, code string, want ast.Program) {
	lexer := lex.NewLexer(code)
	parser := NewParser(lexer)
	tree, err := parser.Parse()
	if !assert.NoError(t, err, name) {
		return
	}

	assert.Equal(t, want, tree, name)
}

func fmtTestName(name, code string) string {
	return fmt.Sprintf("%s (%s)", name, code)
}
