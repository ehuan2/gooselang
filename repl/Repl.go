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
		asts := tokenizer.Parse(line)

		for _, ast := range asts {
			ast.PrintAST()
		}

		if line == "FLY" {
			return
		}

	}

}