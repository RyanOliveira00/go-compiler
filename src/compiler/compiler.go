// src/compiler/compiler.go
package compiler

import (
	"fmt"
	"strconv"

	"github.com/RyanOliveira00/go-compiler/src/ast"
	"github.com/RyanOliveira00/go-compiler/src/lexer"
)

type ValueType int

const (
	ValueTypeInt ValueType = iota
	ValueTypeFloat
	ValueTypeString
	ValueTypeBool
)

type Value struct {
	Type  ValueType
	Value interface{}
}

type Environment struct {
	variables map[string]Value
}

type Compiler struct {
	env *Environment
}

func New() *Compiler {
	return &Compiler{
		env: &Environment{
			variables: make(map[string]Value),
		},
	}
}

func (c *Compiler) Compile(program ast.BlockStmt) (interface{}, error) {
	var result interface{}
	var err error

	for _, stmt := range program.Body {
		result, err = c.executeStmt(stmt)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (c *Compiler) executeStmt(stmt ast.Stmt) (interface{}, error) {
	switch s := stmt.(type) {
	case ast.ExprStmt:
		return c.executeExpr(s.Expression)
	case ast.VarDeclStmt:
		return c.executeVarDecl(s)
	case ast.IfStmt:
		return c.executeIf(s)
	case ast.WhileStmt:
		return c.executeWhile(s)
	case ast.PrintStmt:
		return c.executePrint(s)
	case ast.ReadStmt:
		return c.executeRead(s)
	default:
		return nil, fmt.Errorf("unknown statement type: %T", stmt)
	}
}

func (c *Compiler) executeVarDecl(stmt ast.VarDeclStmt) (interface{}, error) {
	varType := ValueTypeFloat
	if stmt.ExplicitType != nil {
		if typeSymbol, ok := stmt.ExplicitType.(ast.SymbolType); ok {
			switch typeSymbol.Name {
			case "string":
				varType = ValueTypeString
			case "int":
				varType = ValueTypeInt
			case "float":
				varType = ValueTypeFloat
			case "bool":
				varType = ValueTypeBool
			default:
				return nil, fmt.Errorf("unknown type: %s", typeSymbol.Name)
			}
		}
	}

	var defaultValue interface{}
	switch varType {
	case ValueTypeInt:
		defaultValue = int64(0)
	case ValueTypeFloat:
		defaultValue = float64(0)
	case ValueTypeString:
		defaultValue = ""
	case ValueTypeBool:
		defaultValue = false
	}

	if stmt.AssignedValue != nil {
		val, err := c.executeExpr(stmt.AssignedValue)
		if err != nil {
			return nil, err
		}
		defaultValue = val
	}

	c.env.variables[stmt.VariableName] = Value{
		Type:  varType,
		Value: defaultValue,
	}

	return nil, nil
}

func (c *Compiler) executeExpr(expr ast.Expr) (interface{}, error) {
	switch e := expr.(type) {
	case ast.NumberExpr:
		return e.Value, nil
	case ast.StringExpr:
		return e.Value, nil
	case ast.SymbolExpr:
		if value, exists := c.env.variables[e.Value]; exists {
			return value.Value, nil
		}
		return nil, fmt.Errorf("undefined variable: %s", e.Value)
	case ast.BinaryExpr:
		return c.executeBinaryExpr(e)
	case ast.AssignmentExpr:
		return c.executeAssignment(e)
	default:
		return nil, fmt.Errorf("unknown expression type: %T", expr)
	}
}

func (c *Compiler) executeBinaryExpr(expr ast.BinaryExpr) (interface{}, error) {
	left, err := c.executeExpr(expr.Left)
	if err != nil {
		return nil, err
	}

	right, err := c.executeExpr(expr.Right)
	if err != nil {
		return nil, err
	}

	if lstr, lok := left.(string); lok {
		if rstr, rok := right.(string); rok {
			if expr.Operator.Kind == lexer.PLUS {
				return lstr + rstr, nil
			}
			return nil, fmt.Errorf("invalid operation for strings")
		}
	}

	leftNum, err := c.toNumber(left)
	if err != nil {
		return nil, err
	}

	rightNum, err := c.toNumber(right)
	if err != nil {
		return nil, err
	}

	switch expr.Operator.Kind {
	case lexer.PLUS:
		return leftNum + rightNum, nil
	case lexer.DASH:
		return leftNum - rightNum, nil
	case lexer.STAR:
		return leftNum * rightNum, nil
	case lexer.SLASH:
		if rightNum == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		return leftNum / rightNum, nil
	case lexer.LESS:
		return leftNum < rightNum, nil
	case lexer.LESS_EQUALS:
		return leftNum <= rightNum, nil
	case lexer.GREATER:
		return leftNum > rightNum, nil
	case lexer.GREATER_EQUALS:
		return leftNum >= rightNum, nil
	case lexer.EQUALS:
		return leftNum == rightNum, nil
	case lexer.NOT_EQUALS:
		return leftNum != rightNum, nil
	default:
		return nil, fmt.Errorf("unknown operator: %v", expr.Operator)
	}
}

func (c *Compiler) executeAssignment(expr ast.AssignmentExpr) (interface{}, error) {
	target, ok := expr.Assigne.(ast.SymbolExpr)
	if !ok {
		return nil, fmt.Errorf("invalid assignment target")
	}

	value, err := c.executeExpr(expr.Value)
	if err != nil {
		return nil, err
	}

	varInfo, exists := c.env.variables[target.Value]
	if !exists {
		return nil, fmt.Errorf("undefined variable: %s", target.Value)
	}

	varInfo.Value = value
	c.env.variables[target.Value] = varInfo
	return value, nil
}

func (c *Compiler) executeIf(stmt ast.IfStmt) (interface{}, error) {
	condition, err := c.executeExpr(stmt.Condition)
	if err != nil {
		return nil, err
	}

	if isTruthy(condition) {
		return c.executeBlock(stmt.Consequence)
	} else if stmt.Alternative != nil {
		return c.executeBlock(*stmt.Alternative)
	}

	return nil, nil
}

func (c *Compiler) executeWhile(stmt ast.WhileStmt) (interface{}, error) {
	var lastValue interface{}

	for {
		condition, err := c.executeExpr(stmt.Condition)
		if err != nil {
			return nil, err
		}

		if !isTruthy(condition) {
			break
		}

		lastValue, err = c.executeBlock(stmt.Body)
		if err != nil {
			return nil, err
		}
	}

	return lastValue, nil
}

func (c *Compiler) executePrint(stmt ast.PrintStmt) (interface{}, error) {
	value, err := c.executeExpr(stmt.Expression)
	if err != nil {
		return nil, err
	}
	fmt.Println(value)
	return nil, nil
}

func (c *Compiler) executeRead(stmt ast.ReadStmt) (interface{}, error) {
	var input string
	fmt.Scanln(&input)

	target, ok := stmt.Target.(ast.SymbolExpr)
	if !ok {
		return nil, fmt.Errorf("invalid read target")
	}

	varInfo, exists := c.env.variables[target.Value]
	if !exists {
		return nil, fmt.Errorf("undefined variable: %s", target.Value)
	}

	return c.convertInput(input, varInfo.Type)
}

func (c *Compiler) executeBlock(block ast.BlockStmt) (interface{}, error) {
	var result interface{}
	var err error

	for _, stmt := range block.Body {
		result, err = c.executeStmt(stmt)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (c *Compiler) toNumber(v interface{}) (float64, error) {
	switch val := v.(type) {
	case int64:
		return float64(val), nil
	case float64:
		return val, nil
	case int:
		return float64(val), nil
	default:
		return 0, fmt.Errorf("cannot convert %v to number", v)
	}
}

func (c *Compiler) convertInput(input string, targetType ValueType) (interface{}, error) {
	var value interface{}
	var err error

	switch targetType {
	case ValueTypeInt:
		value, err = strconv.ParseInt(input, 10, 64)
	case ValueTypeFloat:
		value, err = strconv.ParseFloat(input, 64)
	case ValueTypeString:
		value = input
	case ValueTypeBool:
		value, err = strconv.ParseBool(input)
	default:
		return nil, fmt.Errorf("unsupported type for read")
	}

	if err != nil {
		return nil, fmt.Errorf("invalid input for type")
	}

	return value, nil
}

func isTruthy(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return v
	case int64:
		return v != 0
	case float64:
		return v != 0
	case string:
		return v != ""
	default:
		return false
	}
}
