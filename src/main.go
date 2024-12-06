package main

import (
	"os"

	"github.com/RyanOliveira00/go-compiler/src/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
