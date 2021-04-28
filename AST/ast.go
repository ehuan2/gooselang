package AST

import "fmt"

type ASTType int

const (
	AST_GOOSE   ASTType = 0
	AST_GOSLING ASTType = 1
	AST_VAR     ASTType = 2
	AST_HONK    ASTType = 3
	AST_BAD     ASTType = 4
	AST_FLY     ASTType = 5
)

type AST interface {
	PrintAST()
	GetType() ASTType
}

func (f Gosling) PrintAST() {
	fmt.Printf("Gosling(%s ", f.param)
	f.body.PrintAST()
	fmt.Printf(")")
}
func MakeGosling(body AST, param string) Gosling {
	return Gosling{body: body, param: param}
}
func (g Gosling) GetType() ASTType {
	return AST_GOSLING
}

// an anonymous function
type Gosling struct {
	body  AST
	param string
}

func (g Goose) PrintAST() {
	fmt.Printf("Goose(%s ", g.name)
	g.value.PrintAST()
	fmt.Printf(")")
}
func MakeGoose(name string, value AST) Goose {
	return Goose{name: name, value: value}
}
func (g Goose) GetType() ASTType {
	return AST_GOOSE
}

// a global function
type Goose struct {
	name  string
	value AST
}

func (v Var) PrintAST() {
	fmt.Printf("Var(%s)", v.name)
}
func MakeVar(name string) Var {
	return Var{name: name}
}
func (g Var) GetType() ASTType {
	return AST_VAR
}
type Var struct {
	name string
}

func (a Honk) PrintAST() {
	fmt.Printf("HONK(")
	a.fn.PrintAST()
	fmt.Printf(")(")
	a.arg.PrintAST()
	fmt.Printf(")")
}
func MakeHonk(fn AST, arg AST) Honk {
	return Honk{fn: fn, arg: arg}
}
func (g Honk) GetType() ASTType {
	return AST_HONK
}

// an application of one function on another
type Honk struct {
	fn  AST
	arg AST
}

func (f Fly) PrintAST() {
	fmt.Printf("~FLY~")
}
func MakeFly() Fly {
	return Fly{}
}
func (g Fly) GetType() ASTType {
	return AST_FLY
}

type Fly struct{}

func (b BadAst) PrintAST() {
	fmt.Printf("Bad GooseLang")
}
func (g BadAst) GetType() ASTType {
	return AST_BAD
}
type BadAst struct{}
