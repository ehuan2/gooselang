package AST

import "fmt"

type AST interface {
	PrintAST()
}

func(f Gosling) PrintAST() {
	fmt.Printf("Gosling %s Honk ", f.param)
	f.body.PrintAST()
	fmt.Printf(" honK")
}
func MakeGosling(body AST, param string) Gosling {
	return Gosling{body: body, param: param}
}
// an anonymous function
type Gosling struct {
	body AST
	param string
}

func (g Goose) PrintAST() {
	fmt.Printf("Goose %s ", g.name)
	g.value.PrintAST()
}
func MakeGoose(name string, value AST) Goose {
	return Goose{name: name, value: value}
} 
// a global function
type Goose struct {
	name string
	value AST
}

func (v Var) PrintAST() {
	fmt.Printf(v.name)
}
func MakeVar(name string) Var {
	return Var{name: name}
}
type Var struct {
	name string
}

func (a Honk) PrintAST() {
	a.fn.PrintAST()
	fmt.Printf(" HONK ")
	a.arg.PrintAST()
}
func MakeHonk(fn AST, arg AST) Honk {
	return Honk{fn: fn, arg: arg}
}
// an application of one function on another
type Honk struct {
	fn AST
	arg AST
}