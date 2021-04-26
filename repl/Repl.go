package repl

import (
	"bufio"
	"fmt"
	"gooselang/tokenizer"
	"io"
)

func Repl(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("> ")
		scanned := scanner.Scan()

		if !scanned {
			continue
		}

		line := scanner.Text()
		ast := tokenizer.Parse(line)

		ast.PrintAST()
		fmt.Println()

		if line == "FLY" {
			return
		}

	}

}