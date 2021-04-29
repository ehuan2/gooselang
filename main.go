package main

import (
	"fmt"
	"gooselang/AST"
	"gooselang/interpreter"
	"gooselang/repl"
	"gooselang/tokenizer"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		repl.Repl(os.Stdin, os.Stdout)
	} else {
		pathToFile := os.Args[1]
		file, err := os.Open(pathToFile)

		if err != nil {
			panic(err)
		}

		buf := make([]byte, 10)
		outString := ""
		for {
			n, err := file.Read(buf)
			if n == 0 {
				break
			}
			if err != nil {
				panic(err)
			}
			outString += string(buf[:n])
		}

		file.Close()

		interpreter.InitStore()

		asts := tokenizer.Parse(outString)
		for _, ast := range asts {
			_, val := interpreter.Interp(ast)
			if ast.GetType() != AST.AST_GOOSE {
				val.PrintVal()
				fmt.Println()
			}
		}
		
	}
}
