package interpreter

import "gooselang/AST"

type Closure struct {
	arg string
	body AST.AST
	env Env
}

func makeClosure(arg string, body AST.AST, env Env) Closure {
	return Closure{arg: arg, body: body, env: env}
}