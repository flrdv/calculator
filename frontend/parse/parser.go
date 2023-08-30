package parse

import (
	"calculator/frontend/lex"
	"fmt"
	"strconv"
)

type Parser struct {
	lexer *lex.Lexer
}

func NewParser(lexer *lex.Lexer) *Parser {
	return &Parser{lexer: lexer}
}

func (p *Parser) Parse() (program Program, err error) {
	for !p.lexer.EOF() {
		node, err := p.stmt()
		if err != nil {
			return nil, err
		}

		program = append(program, node)
	}

	return program, nil
}

func (p *Parser) stmt() (Node, error) {
	expr, err := p.expr()
	if err != nil {
		return nil, err
	}

	for {
		lexeme, err := p.lexer.Next()
		if err != nil {
			return nil, err
		}

		if lexeme.Type == lex.EOF {
			break
		}

		switch lexeme.Type {
		case lex.OpPlus, lex.OpMinus:
			right, err := p.expr()
			if err != nil {
				return nil, err
			}

			expr = BinOp{
				Op:    lexeme.Type,
				Left:  expr,
				Right: right,
			}
		default:
			if lexeme.Type.IsOperator() {
				return nil, fmt.Errorf("unexpected operator: %s", lexeme)
			}

			p.lexer.Back()

			return expr, nil
		}
	}

	return expr, nil
}

func (p *Parser) expr() (Node, error) {
	factor, err := p.term()
	if err != nil {
		return nil, err
	}

	for {
		lexeme, err := p.lexer.Next()
		if err != nil {
			return nil, err
		}

		if lexeme.Type == lex.EOF {
			break
		}

		switch lexeme.Type {
		case lex.OpStar, lex.OpSlash:
			right, err := p.term()
			if err != nil {
				return nil, err
			}

			factor = BinOp{
				Op:    lexeme.Type,
				Left:  factor,
				Right: right,
			}
		default:
			p.lexer.Back()

			return factor, nil
		}
	}

	return factor, nil
}

func (p *Parser) term() (Node, error) {
	factor, err := p.factor()
	if err != nil {
		return nil, err
	}

	for {
		lexeme, err := p.lexer.Next()
		if err != nil {
			return nil, err
		}

		if lexeme.Type == lex.EOF {
			break
		}

		switch lexeme.Type {
		case lex.OpCaret:
			right, err := p.term()
			if err != nil {
				return nil, err
			}

			factor = BinOp{
				Op:    lex.OpCaret,
				Left:  factor,
				Right: right,
			}
		default:
			p.lexer.Back()

			return factor, nil
		}
	}

	return factor, nil
}

func (p *Parser) factor() (Node, error) {
	lexeme, err := p.lexer.Next()
	if err != nil {
		return nil, err
	}

	switch lexeme.Type {
	case lex.Number:
		value, err := strconv.ParseInt(lexeme.Value, 10, 64)
		return Number(value), err
	case lex.Id:
		return ID(lexeme.Value), nil
	case lex.UnPlus, lex.UnMinus:
		value, err := p.term()
		return UnOp{
			Op:    lexeme.Type,
			Value: value,
		}, err
	case lex.LParen:
		stmt, err := p.stmt()
		if err != nil {
			return nil, err
		}

		return stmt, p.match(lex.RParen)
	default:
		return nil, fmt.Errorf("unexpected factor: %s", lexeme)
	}
}

func (p *Parser) match(typ lex.LexemeType) error {
	lexeme, err := p.lexer.Next()
	if err != nil {
		return err
	}

	if lexeme.Type != typ {
		return fmt.Errorf("wanted %s, got %s", typ, lexeme)
	}

	return nil
}
