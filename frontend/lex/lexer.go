package lex

import (
	"errors"
	"fmt"
	"strings"
)

type Lexer struct {
	input          string
	previous       Lexeme
	returnPrevious bool
	pos            Position
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:    input,
		previous: Lexeme{Type: Untyped},
	}
}

func (l *Lexer) Next() (Lexeme, error) {
	if l.returnPrevious {
		l.returnPrevious = false

		return l.previous, nil
	}

	l.skipWhitespaces()

	switch typ := l.guessLexemeType(); typ {
	case EOF:
		return l.save(Lexeme{Type: EOF}), nil
	case Untyped:
		return l.save(Lexeme{}), errors.New("unrecognized lexeme: " + untilSpace(l.input))
	case Number:
		value, err := l.parseNumber()
		return l.save(Lexeme{Number, value}), err
	case Id:
		value, err := l.parseId()

		if isKeyword(value) {
			return l.save(Lexeme{Keyword, value}), err
		}

		return l.save(Lexeme{Id, value}), err
	case symbol:
		lexeme, err := l.parseOperator()
		return l.save(lexeme), err
	case LParen:
		return l.save(Lexeme{LParen, l.after(1)}), nil
	case RParen:
		return l.save(Lexeme{RParen, l.after(1)}), nil
	default:
		panic("BUG: guessLexemeType() returned unknown lexeme type")
	}
}

func (l *Lexer) EOF() bool {
	return len(l.input) == 0
}

func (l *Lexer) Back() {
	l.returnPrevious = true
}

func (l *Lexer) save(lexeme Lexeme) Lexeme {
	l.previous = lexeme
	return lexeme
}

func (l *Lexer) skipWhitespaces() {
	for i, char := range l.input {
		switch char {
		case ' ', '\t', '\r':
			l.pos.Char++
		case '\n':
			l.pos.Char = 0
			l.pos.Line++
		default:
			l.input = l.input[i:]
			return
		}
	}

	l.input = ""
}

func (l *Lexer) parseNumber() (string, error) {
	for i := 0; i < len(l.input); i++ {
		if !isInt(l.input[i]) {
			return l.after(i), nil
		}
	}

	return l.after(len(l.input)), nil
}

func (l *Lexer) parseId() (string, error) {
	for i := 0; i < len(l.input); i++ {
		if !isIdentTail(l.input[i]) {
			return l.after(i), nil
		}
	}

	return l.after(len(l.input)), nil
}

func (l *Lexer) parseOperator() (Lexeme, error) {
	for i := 1; i < len(l.input); i++ {
		if !isUnaryPrefix(l.input[:i+1]) {
			sym := l.after(i)
			symType := symbolType(sym)
			if symType == Untyped {
				return Lexeme{}, fmt.Errorf("unknown operator: %s", sym)
			}

			if l.previous.Type.FollowingSymCanBeUnary() {
				unType := symType.AsUnary()
				if unType == Untyped {
					return Lexeme{}, fmt.Errorf("unknown unary: %s", sym)
				}

				return Lexeme{unType, sym}, nil
			}

			return Lexeme{symType, sym}, nil
		}
	}

	return Lexeme{}, fmt.Errorf("incomplete expression: no right operand")
}

func (l *Lexer) guessLexemeType() LexemeType {
	switch {
	case len(l.input) == 0:
		return EOF
	case isInt(l.input[0]):
		return Number
	case isIdent(l.input[0]):
		return Id
	case isSymbolPrefix(l.input[:1]):
		return symbol
	case l.input[0] == '(':
		return LParen
	case l.input[0] == ')':
		return RParen
	}

	return Untyped
}

func (l *Lexer) after(n int) string {
	l.pos.Char += n
	after := l.input[:n]
	l.input = l.input[n:]

	return after
}

func untilSpace(str string) string {
	before, _, _ := strings.Cut(str, " ")
	return before
}
