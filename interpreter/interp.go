package interpreter

import (
	"fmt"
	"gooselang/AST"
	"os"
)

type bouncerType int

const (
	BOUNCER_INTERP bouncerType = 0
	BOUNCER_CONT   bouncerType = 1
	BOUNCER_EXIT   bouncerType = 2
)

var store map[string]Val // used as global variable for storing GOOSE's
var cont Cont // used as global so we don't pass through
var expr AST.AST
var val Val // the only values are closures or variables
var env Env

func InitStore() {
	store = make(map[string]Val)
}

// it searches the env, then the global store for a particular ast, given a certain key
// if it fails, then it returns a var with the string
func lookupGlobal(key string) Val {
	ast, found := env.lookup(key)
	if found {
		return ast
	}

	storeAst, storeFound := store[key]
	if storeFound {
		return storeAst
	}

	return makeVar(AST.MakeVar(key))
}

// trampoline calls either interpHelper or applyCont, which returns the enum for the next type of function
// so it's not really recursive, it enters each function and exits without going any deeper
func trampoline(b bouncerType) {
	for b != BOUNCER_EXIT {
		if b == BOUNCER_INTERP {
			b = interpHelper()
		} else {
			b = applyCont()
		}
	}
}

func interpHelper() bouncerType {

	exprType := expr.GetType()

	switch exprType {
		// ie an App
		case AST.AST_HONK:
			// we go left, we interp the function
			honk := expr.GetHonk()
			expr = honk.GetFn()
			cont = wrapAppL(honk.GetArg(), env, cont)
			return BOUNCER_INTERP
		case AST.AST_GOSLING:
			// if it's a function, we need to apply it to the continuation
			// we update the value as a closure
			// but we make the continuation the same
			gosling := expr.GetGosling()
			val = makeClosure(gosling.GetParam(), gosling.GetBody(), env)
			return BOUNCER_CONT
		case AST.AST_VAR:
			// we apply the continuation based on our lookup function
			// cont stays, we replace val with it's value from lookup
			inVar := expr.GetVar()
			val = lookupGlobal(inVar.GetName())
			if val.getType() == VAL_AST {
				expr = val.getAST().ast
				return BOUNCER_INTERP
			}
			return BOUNCER_CONT
		case AST.AST_FLY:
			// we don't stop execution here, we stop execution when we apply empty continuation to FLY
			val = makeFly(expr.GetFly())
			return BOUNCER_CONT
		default:
			// should be an error ie a badast
			fmt.Println("Bad Gooselang encountered")
			val = BadVal{}
			return BOUNCER_EXIT
	}
}

func Interp(ast AST.AST) (AST.AST, Val) {

	// we always restart these
	expr = ast
	cont = EmptyCont{}
	env = emptyEnv()

	// we cover the goose at the start, goose's must be the first
	if ast.GetType() == AST.AST_GOOSE {
		expr = ast.GetGoose().GetValue()
		key := ast.GetGoose().GetName()
		trampoline(BOUNCER_INTERP)
		store[key] = val
		return AST.MakeGoose(key, expr), val
	}

	trampoline(BOUNCER_INTERP)
	return expr, val
}

func applyCont() bouncerType {

	contType := cont.getType()

	switch contType {
		case APPL:
			// we replace the environment with the previous one
			// we go right, and interpret the right side argument now
			arg := cont.getAst()
			env = cont.getEnv()
			context := cont.getContext()
			val = makeValAST(arg)
			cont = wrapAppR(val, context)
			return BOUNCER_CONT
		case APPR:
			curVal := cont.getVal()

			for curVal.getType() == VAL_AST {
				unwrap := curVal.getAST().ast
				if unwrap.GetType() == AST.AST_GOSLING {
					unwrapGosling := unwrap.GetGosling()
					curVal = makeClosure(unwrapGosling.GetParam(), unwrapGosling.GetBody(), env)
					break
				}
				if unwrap.GetType() != AST.AST_VAR {
					break
				}
				curVal = lookupGlobal(unwrap.GetVar().GetName())
			}

			// must be a closure
			if curVal.getType() != VAL_CLOSURE {
				fmt.Println("\nThis is the next val:")
				curVal.PrintVal()
				fmt.Println("\nBad Gooselang encountered, not a proper function")
				val = BadVal{}
				return BOUNCER_EXIT
			}

			closure := curVal.getClosure()
			expr = closure.body

			fp := closure.arg
			env = closure.env.add(fp, val)
			cont = cont.getContext()

			return BOUNCER_INTERP

		default:
			// ie it's empty
			if val.getType() == VAL_FLY {
				// stop execution, of everything, we applied a fly
				fmt.Println("Fly away little goose, time to leave the nest")
				os.Exit(0)
			}

			// otherwise, we just leave
			return BOUNCER_EXIT
		
	}
}