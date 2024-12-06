// src/repl/repl.go
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/RyanOliveira00/go-compiler/src/compiler"
	"github.com/RyanOliveira00/go-compiler/src/lexer"
	"github.com/RyanOliveira00/go-compiler/src/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	comp := compiler.New()

	fmt.Fprintln(out, "Bem vindo ao compilador de Go!")

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" || line == "quit" {
			return
		}

		tokens := lexer.Tokenize(line)
		ast := parser.Parse(tokens)

		result, err := comp.Compile(ast)
		if err != nil {
			fmt.Fprintf(out, "Error: %s\n", err)
			continue
		}

		if result != nil {
			fmt.Fprintf(out, "%v\n", result)
		}
	}
}
