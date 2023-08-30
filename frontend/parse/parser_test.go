package parse

import (
	"calculator/frontend/lex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	tcs := []struct {
		Expr string
		AST  Node
	}{
		{"42", Number(42)},
		{"abc", ID("abc")},
		{"2+3", BinOp{lex.OpPlus, Number(2), Number(3)}},
		{
			"2*3+4", BinOp{
				lex.OpPlus,
				BinOp{lex.OpStar, Number(2), Number(3)},
				Number(4),
			},
		},
		{
			"2+3*4", BinOp{
				lex.OpPlus,
				Number(2),
				BinOp{lex.OpStar, Number(3), Number(4)},
			},
		},
		{
			"(2+3)*4", BinOp{
				lex.OpStar,
				BinOp{lex.OpPlus, Number(2), Number(3)},
				Number(4),
			},
		},
		{
			"-(+5)", UnOp{
				Op: lex.UnMinus,
				Value: UnOp{
					Op:    lex.UnPlus,
					Value: Number(5),
				},
			},
		},
		{
			"f(x, x+5)", FCall{
				Target: ID("f"),
				Args: []Node{
					ID("x"), BinOp{lex.OpPlus, ID("x"), Number(5)},
				},
			},
		},
	}

	for _, tc := range tcs {
		testParser(t, tc.Expr, []Node{tc.AST})
	}
}

func testParser(t *testing.T, code string, want Program) {
	lexer := lex.NewLexer(code)
	parser := NewParser(lexer)
	ast, err := parser.Parse()
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, want, ast)
}
