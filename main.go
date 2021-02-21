package main

import (
	"log"
	"os"

	"github.com/hiddenbyte/bf-wasm/lexical"
	"github.com/hiddenbyte/bf-wasm/syntax"
	"github.com/hiddenbyte/bf-wasm/wasm"
)

func main() {
	// Lexical analyzer
	tokenized, err := lexical.Tokenize(os.Stdin)
	if err != nil {
		log.Fatalf("lexical: %s", err)
	}

	// Syntax analyzer
	ast := syntax.Parse(tokenized)

	// WASM code generatation
	err = wasm.EmitText(ast, os.Stdout)
	if err != nil {
		log.Fatalf("wasm: %s", err)
	}
}
