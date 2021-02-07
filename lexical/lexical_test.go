package lexical_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/hiddenbyte/bf-wasm/lexical"
	"github.com/hiddenbyte/bf-wasm/quickcheck"
)

// brainfck random code generator
type BrainfckCode string

func generateRandomBrainfckCode() reflect.Value {
	randomCode := BrainfckCode(quickcheck.StringContainingOnly([]rune{'>', '<', '+', '-', '.', ',', '[', ']'}))
	return reflect.ValueOf(randomCode)
}

func isValidToken(token byte) bool {
	return token >= lexical.IncDataPointerToken && token <= lexical.WhileCloseToken
}

func TestTokenize(t *testing.T) {
	quickcheck.AddRandomGenerator(generateRandomBrainfckCode)

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

	quickcheck.Check(t, "single", func(bfCode BrainfckCode) bool {
		return false
	})
}
