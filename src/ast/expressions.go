package ast

import "github.com/RyanOliveira00/go-compiler/src/lexer"

// --------------------
// LITERAL EXPRESSIONS
// --------------------
type NumberExpr struct {
	Value float64
}

func (n NumberExpr) expr() {}

type StringExpr struct {
	Value string
}

func (n StringExpr) expr() {}

type SymbolExpr struct {
	Value string
}

func (n SymbolExpr) expr() {}

// --------------------
// COMPLEX EXPRESSIONS
// --------------------

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (b BinaryExpr) expr() {}

// -2
type PrefixExpr struct {
	Operator  lexer.Token
	RightExpr Expr
}

func (p PrefixExpr) expr() {}

// a = a + 5
// a += 5
// foo.bar += 5
type AssignmentExpr struct {
	Assigne  Expr
	Operator lexer.Token
	Value    Expr
}

func (a AssignmentExpr) expr() {}
