package parser

import (
	"fmt"
	"strconv"

	"github.com/RyanOliveira00/go-compiler/src/ast"
	"github.com/RyanOliveira00/go-compiler/src/lexer"
)

func parser_expr(p *parser, bp binding_power) ast.Expr {
	tokenKind := p.currentTokenKind()
	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("Could not parse expression: %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)

	for bp_lu[p.currentTokenKind()] > bp {
		led_fn, exists := led_lu[p.currentTokenKind()]

		if !exists {
			panic(fmt.Sprintf("Could not parse expression: %s\n", lexer.TokenKindString(p.currentTokenKind())))
		}

		left = led_fn(p, left, bp_lu[p.currentTokenKind()])
	}

	return left
}

func parser_primary_expr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpr{
			Value: number,
		}
	case lexer.STRING:
		return ast.StringExpr{
			Value: p.advance().Value,
		}
	case lexer.IDENTIFIER:
		return ast.SymbolExpr{
			Value: p.advance().Value,
		}
	default:
		panic(fmt.Sprintf("Could not parse primary expression: %s\n", lexer.TokenKindString(p.currentTokenKind())))
	}
}

func parser_binary_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operator := p.advance()
	right := parser_expr(p, bp)
	return ast.BinaryExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func parser_prefix_expr(p *parser) ast.Expr {
	operator := p.advance()
	rhs := parser_expr(p, default_bp)

	return ast.PrefixExpr{
		Operator:  operator,
		RightExpr: rhs,
	}
}

func parser_assigment_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operator := p.advance()
	rhs := parser_expr(p, bp)
	return ast.AssignmentExpr{
		Assigne:  left,
		Operator: operator,
		Value:    rhs,
	}
}

func parser_grouping_expr(p *parser) ast.Expr {
	p.advance() // Consume the open parenthesis
	expr := parser_expr(p, default_bp)
	p.expect(lexer.CLOSE_PAREN)
	return expr
}
