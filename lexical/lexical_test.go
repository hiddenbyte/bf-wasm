package lexical_test

import (
	"strings"
	"testing"

	"github.com/hiddenbyte/bf-wasm/lexical"
	"github.com/hiddenbyte/bf-wasm/quickcheck"
)

type BrainfckCode string

func randBrainfckCode() interface{} {
	return BrainfckCode(quickcheck.StringContainingOnly([]rune{'>', '<', '+', '-', '.', ',', '[', ']'}))
}

func isValidToken(token byte) bool {
	return token >= lexical.IncDataPointerToken && token <= lexical.WhileCloseToken
}

func TestTokenize(t *testing.T) {
	quickcheck.AddRandomGenerator(randBrainfckCode)

	quickcheck.Check(t, "empty", func() bool {
		tokens, err := lexical.Tokenize(strings.NewReader(""))
		if err != nil {
			return false
		}

		return len(tokens) == 0
	})

	quickcheck.Check(t, "ignore", func(a string) bool {
		tokens, err := lexical.Tokenize(strings.NewReader(a))
		if err != nil {
			return false
		}

		for _, token := range tokens {
			if !isValidToken(token) {
				return false
			}
		}
		return true
	})

	quickcheck.Check(t, "len", func(bfCode BrainfckCode) bool {
		tokens, err := lexical.Tokenize(strings.NewReader(string(bfCode)))
		if err != nil {
			return false
		}
		return len(bfCode) == len(tokens)
	})

	quickcheck.Check(t, "token", func(bfCode BrainfckCode) bool {
		tokensByRune := map[rune]byte{'>': lexical.IncDataPointerToken, '<': lexical.DecDataPointerToken, '+': lexical.IncDataToken, '-': lexical.DecDataToken, ',': lexical.InputToken, '.': lexical.OutputToken, '[': lexical.WhileOpenToken, ']': lexical.WhileCloseToken}

		tokens, err := lexical.Tokenize(strings.NewReader(string(bfCode)))
		if err != nil {
			return false
		}

		for i, r := range bfCode {
			if tokensByRune[r] != tokens[i] {
				return false
			}
		}
		return true
	})
}
