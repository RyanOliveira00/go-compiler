package parser

import (
	"github.com/RyanOliveira00/go-compiler/src/ast"
	"github.com/RyanOliveira00/go-compiler/src/lexer"
)

type binding_power int

const (
	default_bp binding_power = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type stmt_handler func(p *parser) ast.Stmt

type nud_handler func(p *parser) ast.Expr

type led_handler func(p *parser, left ast.Expr, bp binding_power) ast.Expr

type stmt_lookup map[lexer.TokenKind]stmt_handler
type nud_lookup map[lexer.TokenKind]nud_handler
type led_lookup map[lexer.TokenKind]led_handler
type bp_lookup map[lexer.TokenKind]binding_power

var bp_lu = bp_lookup{}
var nud_lu = nud_lookup{}
var led_lu = led_lookup{}
var stmt_lu = stmt_lookup{}

func led(kind lexer.TokenKind, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

func nud(kind lexer.TokenKind, nud_fn nud_handler) {
	nud_lu[kind] = nud_fn
}

func stmt(kind lexer.TokenKind, stmt_fn stmt_handler) {
	bp_lu[kind] = default_bp
	stmt_lu[kind] = stmt_fn
}

func createTokenLookups() {

	led(lexer.ASSIGNMENT, assignment, parser_assigment_expr)
	led(lexer.PLUS_EQUALS, assignment, parser_assigment_expr)
	led(lexer.MINUS_EQUALS, assignment, parser_assigment_expr)
	// TODO add *= /= &=

	// Logical
	led(lexer.AND, logical, parser_binary_expr)
	led(lexer.OR, logical, parser_binary_expr)
	led(lexer.DOT_DOT, logical, parser_binary_expr) // Range Operator (..)

	// Relational
	led(lexer.LESS, relational, parser_binary_expr)
	led(lexer.LESS_EQUALS, relational, parser_binary_expr)
	led(lexer.GREATER, relational, parser_binary_expr)
	led(lexer.GREATER_EQUALS, relational, parser_binary_expr)
	led(lexer.EQUALS, relational, parser_binary_expr)
	led(lexer.NOT_EQUALS, relational, parser_binary_expr)

	// Additive & Multiplicative
	led(lexer.PLUS, additive, parser_binary_expr)
	led(lexer.DASH, additive, parser_binary_expr)

	led(lexer.STAR, multiplicative, parser_binary_expr)
	led(lexer.SLASH, multiplicative, parser_binary_expr)
	led(lexer.PERCENT, multiplicative, parser_binary_expr)

	// Literals & Symbols
	nud(lexer.NUMBER, parser_primary_expr)
	nud(lexer.STRING, parser_primary_expr)
	nud(lexer.IDENTIFIER, parser_primary_expr)
	nud(lexer.OPEN_PAREN, parser_grouping_expr)
	nud(lexer.DASH, parser_prefix_expr)

	// Statements
	stmt(lexer.CONST, parser_var_decl_stmt)
	stmt(lexer.LET, parser_var_decl_stmt)
	stmt(lexer.IF, parser_if_stmt)
	stmt(lexer.WHILE, parser_while_stmt)
	stmt(lexer.PRINT, parser_print_stmt)
	stmt(lexer.READ, parser_read_stmt)
}
