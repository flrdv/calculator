package lex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLexer(t *testing.T) {
	t.Run("single digit", func(t *testing.T) {
		testLexer(t, "1", Lexeme{Number, "1"})
	})

	t.Run("single number", func(t *testing.T) {
		testLexer(t, "125", Lexeme{Number, "125"})
	})

	t.Run("single letter", func(t *testing.T) {
		testLexer(t, "a", Lexeme{Id, "a"})
	})

	t.Run("single identifier", func(t *testing.T) {
		testLexer(t, "abc", Lexeme{Id, "abc"})
	})

	t.Run("simple expression", func(t *testing.T) {
		testLexer(
			t, "a+5",
			Lexeme{Id, "a"}, Lexeme{OpPlus, "+"}, Lexeme{Number, "5"},
		)
	})

	t.Run("no right operand", func(t *testing.T) {
		lexer := NewLexer("5+")
		_, err := lexer.Next()
		require.NoError(t, err)
		_, err = lexer.Next()
		require.EqualError(t, err, "incomplete expression: no right operand")
	})

	t.Run("unary", func(t *testing.T) {
		testLexer(
			t, "+5+-+-7",
			Lexeme{UnPlus, "+"}, Lexeme{Number, "5"},
			Lexeme{OpPlus, "+"}, Lexeme{UnMinus, "-"},
			Lexeme{UnPlus, "+"}, Lexeme{UnMinus, "-"},
			Lexeme{Number, "7"},
		)
	})

	t.Run("comma", func(t *testing.T) {
		testLexer(t, "a,b", Lexeme{Id, "a"}, Lexeme{ChComma, ","}, Lexeme{Id, "b"})
	})
}

func testLexer(t *testing.T, code string, want ...Lexeme) {
	lexer := NewLexer(code)

	for _, wantedLexeme := range want {
		lexeme, err := lexer.Next()
		require.NoError(t, err)
		require.Equal(t, wantedLexeme, lexeme)
	}

	lexeme, err := lexer.Next()
	require.NoError(t, err)
	require.Equal(t, Lexeme{Type: EOF}, lexeme)
}
