package interpreter

import (
	"fmt"
	"gooselang/AST"
)

// an environment is really just a map from string keys to asts
type Env struct {
	env map[string]AST.AST
}

func emptyEnv() Env {
	newEnv := make(map[string]AST.AST)
	return Env{env: newEnv}
}

func (e Env) printEnv() {
	fmt.Print(e.env)
}

func (e Env) lookup(key string) AST.AST {
	return e.env[key]
}

func (e Env) add(key string, value AST.AST) Env {
	e.env[key] = value
	return e
}