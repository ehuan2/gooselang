package interpreter

import (
	"fmt"
	"gooselang/AST"
)

type ValType int

const (
	VAL_CLOSURE ValType = 0
	VAL_VAR     ValType = 1
	VAL_FLY			ValType = 2
	VAL_BAD			ValType = 3
	VAL_AST			ValType = 4
)

// values are either vars or closures
type Val interface {
	getType() ValType
	getClosure() Closure
	getVar() Var
	getAST() ValAST
	PrintVal()
}

func (nc notClosure) getClosure() Closure {
	return Closure{}
}
func (nc notVar) getVar() Var {
	return Var{}
}
func (nc notAST) getAST() ValAST {
	return ValAST{}
}
type notClosure struct {}
type notVar struct {}
type notAST struct {}

func (c Closure) getType() ValType {
	return VAL_CLOSURE
}

func (c Closure) getClosure() Closure {
	return c
}
func (c Closure) PrintVal() {
	fmt.Printf("Closure(%s ", c.arg)
	c.body.PrintAST()
	fmt.Printf(")")
}

type Closure struct {
	arg  string
	body AST.AST
	env  Env
	notVar
	notAST
}

func makeClosure(arg string, body AST.AST, env Env) Closure {
	return Closure{arg: arg, body: body, env: env}
}

func (v Var) getType() ValType {
	return VAL_VAR
}
func makeVar(v AST.Var) Var {
	return Var{astVar: v}
}
func (v Var) getVar() Var {
	return v
}
func (v Var) PrintVal() {
	v.astVar.PrintAST()
}
type Var struct {
	astVar AST.Var
	notClosure
	notAST
}

func (v ValAST) getType() ValType {
	return VAL_AST
}
func makeValAST(v AST.AST) ValAST {
	return ValAST{ast: v}
}
func (v ValAST) getAST() ValAST {
	return v
}
func (v ValAST) PrintVal() {
	fmt.Printf("AST(")
	v.ast.PrintAST()
	fmt.Printf(")")
}
type ValAST struct {
	ast AST.AST
	notClosure
	notVar
}

func (f Fly) getType() ValType {
	return VAL_FLY
}
func makeFly(f AST.Fly) Fly {
	return Fly{fly: f}
}
func (f Fly) PrintVal() {
	fmt.Printf("~ValFLY~")
}
type Fly struct {
	fly AST.Fly
	notClosure
	notVar
	notAST
}

func (v BadVal) getType() ValType {
	return VAL_BAD
}
func (v BadVal) PrintVal() {
	fmt.Printf("Bad Value")
}
type BadVal struct {
	notClosure
	notVar
	notAST
}