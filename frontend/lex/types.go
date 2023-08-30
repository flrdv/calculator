package lex

import "fmt"

type LexemeType string

const (
	EOF             LexemeType = "EOF"
	Untyped         LexemeType = ""
	Number          LexemeType = "NUMBER"
	untypedOperator LexemeType = "OPERATOR"
	OpPlus          LexemeType = "OP_PLUS"
	OpMinus         LexemeType = "OP_MINUS"
	OpStar          LexemeType = "OP_STAR"
	OpSlash         LexemeType = "OP_SLASH"
	OpCaret         LexemeType = "OP_CARET"
	UnPlus          LexemeType = "UN_PLUS"
	UnMinus         LexemeType = "UN_MINUS"
	Id              LexemeType = "ID"
	Keyword         LexemeType = "KEYWORD"
	LParen          LexemeType = "LPAREN"
	RParen          LexemeType = "RPAREN"
)

func (l LexemeType) IsOperator() bool {
	switch l {
	case untypedOperator, OpPlus, OpMinus, OpStar, OpSlash, OpCaret, UnPlus, UnMinus:
		return true
	}

	return false
}

func (l LexemeType) FollowingOpCanBeUnary() bool {
	return l.IsOperator() || l == Untyped || l == LParen
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
