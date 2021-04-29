package tokenizer

import (
	"fmt"
	"gooselang/AST"
	"strings"
)

func isImproperVar(id string) bool {
	return id == "Honk" || id == "honK" || id == "FLY" || id == "Goose" || id == "Gosling" || id == "HONK"
}

func doParse(sexp SExp) (AST.AST) {

	// if it's bad, return bad right away
	if sexp.getType() == BADSEXP {
		return AST.BadAst{}
	}

	// for atoms, it can start with Goose, Gosling, HONK, a variable or FLY (DONE)
	// but we only care about variables and FLY, other cases are dealt within lists
	if sexp.getType() == ATOM {
		token := sexp.getAtom().atom

		if token.tokenType == ID {
			// in case it's a keyword
			if isImproperVar(token.value) {
				return AST.BadAst{}
			}

			// otherwise, return a var
			return AST.MakeVar(token.value)

		} else if token.tokenType == DONE {
			return AST.MakeFly()
		}

	} else if sexp.getType() == LIST {

		list := sexp.getSExpListNode()

		if isBadList(list) {
			return AST.BadAst{}
		}

		length := length(list)

		// the length must be 3, always includes Gosling/Goose/Honk, var-name/argument and then the body
		if length != 3 {
			return AST.BadAst{}
		}

		// the first one must be an atom
		if list.first.getType() != ATOM {
			return AST.BadAst{}
		}

		atom := list.first.getAtom().atom
		atomType := atom.tokenType

		if atomType != HONK {
			// ie it's only Gosling or Goose
			if list.rest.first.getType() != ATOM {
				return AST.BadAst{}
			}

			secondAtom := list.rest.first.getAtom().atom

			if secondAtom.tokenType != ID {
				return AST.BadAst{}
			}

			varName := secondAtom.value

			body := doParse(list.rest.rest.first)

			// if the body is falsy, return a falsy one
			if body.GetType() == AST.AST_BAD {
				return AST.BadAst{}
			}

			if atomType == GOOSE {
				return AST.MakeGoose(varName, body)
			} else if atomType == GOSLING {
				return AST.MakeGosling(body, varName)
			}

			// if not a goose nor gosling, something went wrong
			return AST.BadAst{}

		} else {
			// in case that it's a HONK
			arg := doParse(list.rest.first)
			fnSExp := list.rest.rest.first
			fn := doParse(fnSExp)
			return AST.MakeHonk(fn, arg)

		}
		
	}

	// in case it's not a list or atom, something went horribly wrong
	return AST.BadAst{}
}

func Parse(files string) []AST.AST {
	words := strings.Fields(files)
	tokens := tokenize(words)
	sexps := parseSExp(tokens)

	length := len(sexps)
	astCounter := 0
	out := make([]AST.AST, length)

	for _, sexp := range sexps {
		printSExp(sexp)
		fmt.Println()
	}

	for i := 0; i < length; i++ {
		out[astCounter] = doParse(sexps[i])
		astCounter++
	}

	// only include slots that actually are used up
	return out
}