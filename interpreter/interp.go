package interpreter

import "gooselang/AST"

type bouncerType int

const (
	BOUNCER_INTERP bouncerType = 0
	BOUNCER_CONT   bouncerType = 1
	BOUNCER_EXIT   bouncerType = 2
)

var store map[string]AST.AST // used as global variable for storing GOOSE's
var cont Cont // used as global so we don't pass through
var expr AST.AST

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
			
			return BOUNCER_INTERP
		
	}


	return BOUNCER_EXIT
}

func Interp(ast AST.AST) AST.AST {

	// we always restart these
	expr = ast
	cont = EmptyCont{}

	trampoline(BOUNCER_INTERP)
	return expr
}

func applyCont() bouncerType {

	return BOUNCER_CONT
}
