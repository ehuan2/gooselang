package tokenizer

import (
	"gooselang/AST"
	"strings"
)

func Parse(files string) AST.AST {
	words := strings.Fields(files)
	tokens := Tokenize(words)
	sexps := ParseSExp(tokens)

	return AST.MakeVar("x")
}