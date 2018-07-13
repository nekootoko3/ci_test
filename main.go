package main

import (
	"fmt"
	"os"

	"app/repl"
)

const PROMPT = ">> "

func main() {
	fmt.Print("This is Easy Calculator\n")
	repl.Start(os.Stdin, os.Stdout)
}
