package interpreter

import (
	"fmt"
	"gooselang/AST"
)

type ContType int

const (
	EMPTY ContType = 0
	APPL  ContType = 1
	APPR  ContType = 2
)

// we have two things we can do with our continuation, either go left or right in an application
// or it's empty, cont = MT | AppL Ast Env Cont | AppR Val Cont
// where Val = Closure String Ast Env
type Cont interface {
	getAst() AST.AST
	getEnv() Env
	getVal() AST.AST
	getType() ContType
	getContext() Cont
}

// these shouldn't actually be used
func (a notAppL) getAst() AST.AST {
	return AST.BadAst{}
}
func (a notAppL) getEnv() Env {
	return emptyEnv()
}
type notAppL struct {}

func (a notAppR) getVal() AST.AST {
	return AST.BadAst{}
}
type notAppR struct {}

func (e EmptyCont) getType() ContType {
	return EMPTY
}
// hopefully the following won't get used
func (e EmptyCont) getContext() Cont {
	return e
}
type EmptyCont struct {
	notAppL
	notAppR
}

func (a AppL) getAst() AST.AST {
	return a.ast
}
func (a AppL) getEnv() Env {
	return a.env
}
func (a AppL) getType() ContType {
	return APPL
}
func (a AppL) getContext() Cont {
	return a.context
}
type AppL struct {
	notAppR
	ast AST.AST
	env Env
	context Cont
}

func (a AppR) getType() ContType {
	return APPR
}
func (a AppR) getVal() AST.AST {
	return a.val
}
func (a AppR) getContext() Cont {
	return a.context
}
type AppR struct {
	notAppL
	val AST.AST
	context Cont
}

func wrapAppL(ast AST.AST, env Env, context Cont) Cont {
	return AppL{ast: ast, env: env, context: context}
}

func wrapAppR(val AST.AST, context Cont) Cont {
	return AppR{val: val, context: context}
}

func printCont(cont Cont) {
	contType := cont.getType()
	switch contType {
		case APPL:
			fmt.Printf("(AppL ")
			cont.getAst().PrintAST()
			fmt.Printf(" ")
			cont.getEnv().printEnv()
			fmt.Printf(" ")
			printCont(cont.getContext())
			fmt.Printf(")")
		case APPR:
			fmt.Printf("(AppR ")
			cont.getVal().PrintAST()
			fmt.Printf(" ")
			printCont(cont.getContext())
			fmt.Printf(")")
		default:
			fmt.Printf("MT")
	}
}