package lex

import "strings"

var (
	allSymbols   = []string{plus, minus, star, slash, caret, comma, equal, flow}
	unarySymbols = []string{plus, minus}
)

const (
	plus  = "+"
	minus = "-"
	star  = "*"
	slash = "/"
	caret = "^"
	comma = ","
	equal = "="
	flow  = "->"
)

func symbolType(o string) LexemeType {
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
	case comma:
		return ChComma
	case equal:
		return ChEqual
	case flow:
		return ChFlow
	}

	return Untyped
}

func isSymbolPrefix(opPrefix string) bool {
	for _, op := range allSymbols {
		if strings.HasPrefix(op, opPrefix) {
			return true
		}
	}

	return false
}

func isUnaryPrefix(prefix string) bool {
	for _, op := range unarySymbols {
		if strings.HasPrefix(op, prefix) {
			return true
		}
	}

	return false
}
