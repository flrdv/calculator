package parse

import (
	"calculator/frontend/lex"
	"calculator/frontend/operator"
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

		switch lexeme.Type {
		case lex.EOF:
		case lex.Operator:
			switch op := operator.Operator(lexeme.Value); op {
			case operator.Plus, operator.Minus:
				right, err := p.stmt()

				return BinOp{
					Op:    op,
					Left:  expr,
					Right: right,
				}, err
			default:
				return nil, fmt.Errorf("unexpected operator: (%s %s)", lexeme.Type, lexeme.Value)
			}
		default:
			return nil, fmt.Errorf("unexpected stmt: (%s %s)", lexeme.Type, lexeme.Value)
		}

		return expr, nil
	}
}

func (p *Parser) expr() (Node, error) {
	factor, err := p.factor()
	if err != nil {
		return nil, err
	}

	for {
		lexeme, err := p.lexer.Next()
		if err != nil {
			return nil, err
		}

		switch lexeme.Type {
		case lex.EOF:
		case lex.Operator:
			switch op := operator.Operator(lexeme.Value); op {
			case operator.Star, operator.Slash:
				right, err := p.expr()

				return BinOp{
					Op:    op,
					Left:  factor,
					Right: right,
				}, err
			default:
				p.lexer.Back()

				return factor, nil
			}
		default:
			return nil, fmt.Errorf("unexpected expr: (%s %s)", lexeme.Type, lexeme.Value)
		}

		return factor, nil
	}
}

func (p *Parser) factor() (Node, error) {
	lexeme, err := p.lexer.Next()
	if err != nil {
		return nil, err
	}

	switch lexeme.Type {
	case lex.Number:
		return strconv.Atoi(lexeme.Value)
	case lex.Id:
		return lexeme.Value, nil
	case lex.LParen:
		stmt, err := p.stmt()
		if err != nil {
			return nil, err
		}

		return stmt, p.match(lex.LParen)
	default:
		return nil, fmt.Errorf("unexpected factor: (%s %s)", lexeme.Type, lexeme.Value)
	}
}

func (p *Parser) match(typ lex.LexemeType) error {
	lexeme, err := p.lexer.Next()
	if err != nil {
		return err
	}

	if lexeme.Type != typ {
		return fmt.Errorf("wanted %s, got (%s %s)", typ, lexeme.Type, lexeme.Value)
	}

	return nil
}
