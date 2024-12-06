package parser

import (
	"fmt"

	"github.com/RyanOliveira00/go-compiler/src/ast"
	"github.com/RyanOliveira00/go-compiler/src/lexer"
)

type type_nud_handler func(p *parser) ast.Type
type type_led_handler func(p *parser, left ast.Type, bp binding_power) ast.Type

type type_nud_lookup map[lexer.TokenKind]type_nud_handler
type type_led_lookup map[lexer.TokenKind]type_led_handler
type type_bp_lookup map[lexer.TokenKind]binding_power

var type_bp_lu = type_bp_lookup{}
var type_nud_lu = type_nud_lookup{}
var type_led_lu = type_led_lookup{}

func type_led(kind lexer.TokenKind, bp binding_power, led_fn type_led_handler) {
	type_bp_lu[kind] = bp
	type_led_lu[kind] = led_fn
}

func type_nud(kind lexer.TokenKind, nud_fn type_nud_handler) {
	type_nud_lu[kind] = nud_fn
}

func createTypeLookups() {
	type_nud(lexer.IDENTIFIER, parse_symbol_type)
	type_nud(lexer.OPEN_BRACKET, parse_array_type)
}

func parse_symbol_type(p *parser) ast.Type {
	return ast.SymbolType{
		Name: p.expect(lexer.IDENTIFIER).Value,
	}
}

func parse_array_type(p *parser) ast.Type {
	p.advance()
	p.expect(lexer.CLOSE_BRACKET)
	var underlyingType = parser_type(p, default_bp)
	return ast.ArrayType{
		Underlying: underlyingType,
	}
}

func parser_type(p *parser, bp binding_power) ast.Type {
	tokenKind := p.currentTokenKind()
	nud_fn, exists := type_nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("Could not parse type expression: %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)

	for type_bp_lu[p.currentTokenKind()] > bp {
		led_fn, exists := type_led_lu[p.currentTokenKind()]

		if !exists {
			panic(fmt.Sprintf("Could not parse expression: %s\n", lexer.TokenKindString(p.currentTokenKind())))
		}

		left = led_fn(p, left, type_bp_lu[p.currentTokenKind()])
	}

	return left
}
