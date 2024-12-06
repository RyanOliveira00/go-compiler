package ast

type SymbolType struct {
	Name string // T
}

func (t SymbolType) _type() {}

type ArrayType struct {
	Underlying Type // []T
}

func (t ArrayType) _type() {}
