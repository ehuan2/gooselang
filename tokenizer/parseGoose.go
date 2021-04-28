package tokenizer

import (
	"fmt"
	"gooselang/AST"
	"strings"
)

func isImproperVar(id string) bool {
	return id == "Honk" || id == "honK" || id == "FLY" || id == "Goose" || id == "Gosling"
}

func doParse(sexp SExp, position int) (AST.AST, int) {

	// if it's bad, return bad right away
	if sexp.getType() == BADSEXP {
		return AST.BadAst{}, position + 1
	}

	if sexp.getType() == ATOM {
		token := sexp.getAtom().atom

		if token.tokenType == ID {
			// in case it's a keyword
			if isImproperVar(token.value) {
				return AST.BadAst{}, position + 1
			}

			

		}

	}



	return nil, position + 1
}

func Parse(files string) []AST.AST {
	words := strings.Fields(files)
	tokens := Tokenize(words)
	sexps := ParseSExp(tokens)

	length := len(sexps)
	astCounter := 0
	out := make([]AST.AST, length)

	for _, sexp := range sexps {
		PrintSExp(sexp)
		fmt.Println()
	}

	for i := 0; i < length; {
		out[astCounter], i = doParse(sexps[i], i)
		astCounter++
	}

	// only include slots that actually are used up
	return nil
}