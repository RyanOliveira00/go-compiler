package parser

import (
	"github.com/RyanOliveira00/go-compiler/src/ast"
	"github.com/RyanOliveira00/go-compiler/src/lexer"
)

func parser_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return stmt_fn(p)
	}

	expression := parser_expr(p, default_bp)

	p.expect(lexer.SEMI_COLON)

	return ast.ExprStmt{
		Expression: expression,
	}
}

func parser_if_stmt(p *parser) ast.Stmt {
	p.advance()

	p.expect(lexer.OPEN_PAREN)
	condition := parser_expr(p, default_bp)
	p.expect(lexer.CLOSE_PAREN)

	p.expect(lexer.OPEN_CURLY)
	var consequenceStmts []ast.Stmt
	for p.currentTokenKind() != lexer.CLOSE_CURLY {
		stmt := parser_stmt(p)
		consequenceStmts = append(consequenceStmts, stmt)
	}
	p.expect(lexer.CLOSE_CURLY)
	consequence := ast.BlockStmt{Body: consequenceStmts}

	var alternative *ast.BlockStmt
	if p.currentTokenKind() == lexer.ELSE {
		p.advance()
		p.expect(lexer.OPEN_CURLY)
		var alternativeStmts []ast.Stmt
		for p.currentTokenKind() != lexer.CLOSE_CURLY {
			stmt := parser_stmt(p)
			alternativeStmts = append(alternativeStmts, stmt)
		}
		p.expect(lexer.CLOSE_CURLY)
		alt := ast.BlockStmt{Body: alternativeStmts}
		alternative = &alt
	}

	p.expect(lexer.SEMI_COLON)

	return ast.IfStmt{
		Condition:   condition,
		Consequence: consequence,
		Alternative: alternative,
	}
}

func parser_block_stmt(p *parser) ast.BlockStmt {
	p.expect(lexer.OPEN_CURLY)
	var statements []ast.Stmt

	for p.currentTokenKind() != lexer.CLOSE_CURLY {
		if p.currentTokenKind() == lexer.EOF {
			panic("Unexpected end of file while parsing block")
		}
		stmt := parser_stmt(p)
		statements = append(statements, stmt)
	}

	p.expect(lexer.CLOSE_CURLY)

	return ast.BlockStmt{
		Body: statements,
	}
}

func parser_while_stmt(p *parser) ast.Stmt {
	p.advance()
	p.expect(lexer.OPEN_PAREN)
	condition := parser_expr(p, default_bp)
	p.expect(lexer.CLOSE_PAREN)

	body := parser_block_stmt(p)

	return ast.WhileStmt{
		Condition: condition,
		Body:      body,
	}
}

func parser_print_stmt(p *parser) ast.Stmt {
	p.advance()
	p.expect(lexer.OPEN_PAREN)
	expr := parser_expr(p, default_bp)
	p.expect(lexer.CLOSE_PAREN)
	p.expect(lexer.SEMI_COLON)

	return ast.PrintStmt{
		Expression: expr,
	}
}

func parser_read_stmt(p *parser) ast.Stmt {
	p.advance()
	p.expect(lexer.OPEN_PAREN)
	target := parser_expr(p, default_bp)
	p.expect(lexer.CLOSE_PAREN)
	p.expect(lexer.SEMI_COLON)

	return ast.ReadStmt{
		Target: target,
	}
}

func parser_function_stmt(p *parser) ast.Stmt {
	p.advance()
	name := p.expect(lexer.IDENTIFIER).Value

	p.expect(lexer.OPEN_PAREN)
	parameters := []string{}

	if p.currentTokenKind() != lexer.CLOSE_PAREN {
		for {
			param := p.expect(lexer.IDENTIFIER).Value
			parameters = append(parameters, param)

			if p.currentTokenKind() != lexer.COMMA {
				break
			}
			p.advance()
		}
	}
	p.expect(lexer.CLOSE_PAREN)

	var returnType ast.Type
	if p.currentTokenKind() == lexer.COLON {
		p.advance()
		returnType = parser_type(p, default_bp)
	}

	body := parser_block_stmt(p)

	return ast.FunctionDeclStmt{
		Name:       name,
		Parameters: parameters,
		ReturnType: returnType,
		Body:       body,
	}
}

func parser_return_stmt(p *parser) ast.Stmt {
	p.advance()

	var value ast.Expr
	if p.currentTokenKind() != lexer.SEMI_COLON {
		value = parser_expr(p, default_bp)
	}
	p.expect(lexer.SEMI_COLON)

	return ast.ReturnStmt{
		Value: value,
	}
}

func parser_var_decl_stmt(p *parser) ast.Stmt {
	var explicitType ast.Type
	var assignedValue ast.Expr
	isConst := p.advance().Kind == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "Expected identifier").Value

	if p.currentTokenKind() == lexer.COLON {
		p.advance()
		explicitType = parser_type(p, default_bp)
	}

	if p.currentTokenKind() != lexer.SEMI_COLON {
		p.expect(lexer.ASSIGNMENT)
		assignedValue = parser_expr(p, assignment)
	} else if explicitType == nil {
		panic("Cannot declare variable without an explicit type")
	}

	p.expect(lexer.SEMI_COLON)

	if isConst && assignedValue == nil {
		panic("Cannot declare constant without an assigned")
	}

	return ast.VarDeclStmt{
		ExplicitType:  explicitType,
		IsConstant:    isConst,
		VariableName:  varName,
		AssignedValue: assignedValue,
	}
}
