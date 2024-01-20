package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/miselin/c64lsp/pkg/grammar"
)

func main() {
	code, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read source file: %v", err)
		os.Exit(1)
	}

	g := grammar.NewGrammar()
	p, err := g.Parse(os.Args[1], string(code))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse source file: %v", err)
		os.Exit(1)
	}

	// p.Dump()

	line := p.FindTextLine(9)
	grammar.DumpLine(line)

	tok := p.FindBasicTokenAt(9, 35)
	if tok != nil {
		fmt.Printf("%s\n", *tok)
	} else {
		fmt.Printf("no\n")
	}
}
