package lex

import "fmt"

type LexemeType string

const (
	Untyped LexemeType = ""
	EOF     LexemeType = "EOF"
	Number  LexemeType = "NUMBER"
	symbol  LexemeType = "SYMBOL"
	OpPlus  LexemeType = "OP_PLUS"
	OpMinus LexemeType = "OP_MINUS"
	OpStar  LexemeType = "OP_STAR"
	OpSlash LexemeType = "OP_SLASH"
	OpCaret LexemeType = "OP_CARET"
	UnPlus  LexemeType = "UN_PLUS"
	UnMinus LexemeType = "UN_MINUS"
	ChComma LexemeType = "CH_COMMA"
	ChEqual LexemeType = "CH_EQUAL"
	ChFlow  LexemeType = "CH_FLOW"
	Id      LexemeType = "ID"
	Keyword LexemeType = "KEYWORD"
	LParen  LexemeType = "LPAREN"
	RParen  LexemeType = "RPAREN"
)

func (l LexemeType) IsSymbol() bool {
	switch l {
	case symbol, OpPlus, OpMinus, OpStar, OpSlash, OpCaret, UnPlus, UnMinus:
		return true
	}

	return false
}

func (l LexemeType) FollowingSymCanBeUnary() bool {
	return l.IsSymbol() || l == Untyped || l == LParen
}

func (l LexemeType) AsUnary() LexemeType {
	switch l {
	case OpPlus:
		return UnPlus
	case OpMinus:
		return UnMinus
	}

	return Untyped
}

type Lexeme struct {
	Type  LexemeType
	Value string
}

func (l Lexeme) String() string {
	if len(l.Value) == 0 {
		return fmt.Sprintf("(%s)", l.Type)
	}

	return fmt.Sprintf("(%s %s)", l.Type, l.Value)
}

type Position struct {
	Line, Char int
}
