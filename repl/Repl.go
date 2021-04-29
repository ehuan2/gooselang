package repl

import (
	"bufio"
	"fmt"
	"gooselang/AST"
	"gooselang/interpreter"
	"gooselang/tokenizer"
	"io"
)

func Repl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	interpreter.InitStore()

	for {
		fmt.Printf("> ")
		scanned := scanner.Scan()

		if !scanned {
			continue
		}

		line := scanner.Text()
		asts := tokenizer.Parse(line)

		for _, ast := range asts {
			_, val := interpreter.Interp(ast)
			if ast.GetType() != AST.AST_GOOSE {
				val.PrintVal()
				fmt.Println()
			}
		// 	ast.PrintAST()
		// 	fmt.Println()
		// 	out, val := interpreter.Interp(ast)
		// 	fmt.Printf("Out: ")
		// 	out.PrintAST()
		// 	fmt.Println()
		// 	fmt.Printf("Val: ")
		// 	val.PrintVal()
		// 	fmt.Println()
		}

	}

}