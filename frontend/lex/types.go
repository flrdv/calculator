package lex

type LexemeType string

const (
	EOF      LexemeType = "EOF"
	Untyped  LexemeType = "UNTYPED"
	Number   LexemeType = "NUMBER"
	Operator LexemeType = "OPERATOR"
	Unary    LexemeType = "UNARY"
	Id       LexemeType = "ID"
	Keyword  LexemeType = "KEYWORD"
	LParen   LexemeType = "LPAREN"
	RParen   LexemeType = "RPAREN"
)

type Lexeme struct {
	Type  LexemeType
	Value string
}

type Position struct {
	Line, Char int
}
