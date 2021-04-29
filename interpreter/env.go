package interpreter

import (
	"fmt"
)

// an environment is really just a map from string keys to asts
type Env struct {
	env map[string]Val
}

func emptyEnv() Env {
	newEnv := make(map[string]Val)
	return Env{env: newEnv}
}

func (e Env) printEnv() {
	fmt.Print(e.env)
}

func (e Env) lookup(key string) (Val, bool) {
	ast, found := e.env[key]
	return ast, found
}

func (e Env) add(key string, value Val) Env {
	e.env[key] = value
	return e
}