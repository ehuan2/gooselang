package main

import (
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

		ast := tokenizer.Parse(outString)
		ast.PrintAST()
		
	}
}
