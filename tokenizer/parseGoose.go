package tokenizer

import (
	"fmt"
	"gooselang/AST"
	"strings"
)

// func isImproperVar(id string) bool {
// 	return id == "Honk" || id == "honK" || id == "FLY" || id == "Goose" || id == "Gosling"
// }

func Parse(files string) []AST.AST {
	words := strings.Fields(files)
	tokens := Tokenize(words)
	sexps := ParseSExp(tokens)
	for _, sexp := range sexps {
		PrintSExp(sexp)
		fmt.Println()
	}
	return nil
}