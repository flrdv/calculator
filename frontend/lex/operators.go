package lex

import "strings"

var (
	allOps = []string{plus, minus, star, slash, caret}
	unOps  = []string{plus, minus}
)

const (
	plus  = "+"
	minus = "-"
	star  = "*"
	slash = "/"
	caret = "^"
)

func operatorType(o string) LexemeType {
	switch o {
	case plus:
		return OpPlus
	case minus:
		return OpMinus
	case star:
		return OpStar
	case slash:
		return OpSlash
	case caret:
		return OpCaret
	}

	return Untyped
}

func isOperatorPrefix(opPrefix string) bool {
	for _, op := range allOps {
		if strings.HasPrefix(op, opPrefix) {
			return true
		}
	}

	return false
}

func isUnaryPrefix(prefix string) bool {
	for _, op := range unOps {
		if strings.HasPrefix(op, prefix) {
			return true
		}
	}

	return false
}
