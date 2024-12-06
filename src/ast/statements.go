package ast

// { ... []Stmt }

type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt() {}

type ExprStmt struct {
	Expression Expr
}

func (e ExprStmt) stmt() {}

type VarDeclStmt struct {
	VariableName  string
	IsConstant    bool
	AssignedValue Expr
	ExplicitType  Type
}

func (e VarDeclStmt) stmt() {}

type IfStmt struct {
	Condition   Expr
	Consequence BlockStmt
	Alternative *BlockStmt
}

func (i IfStmt) stmt() {}

type WhileStmt struct {
	Condition Expr
	Body      BlockStmt
}

func (w WhileStmt) stmt() {}

type PrintStmt struct {
	Expression Expr
}

func (p PrintStmt) stmt() {}

type ReadStmt struct {
	Target Expr
}

func (r ReadStmt) stmt() {}

type FunctionDeclStmt struct {
	Name       string
	Parameters []string
	ReturnType Type
	Body       BlockStmt
}

func (f FunctionDeclStmt) stmt() {}

type ReturnStmt struct {
	Value Expr
}

func (r ReturnStmt) stmt() {}
