package lexical_test

import (
	"strings"
	"testing"

	"github.com/hiddenbyte/bf-wasm/lexical"
	"github.com/hiddenbyte/bf-wasm/quickcheck"
)

// GenBrainFckWithoutComments generates lexically valid brainfck code  without any comments
func GenBrainFckWithoutComments() string {
	return quickcheck.StringContainingOnly([]rune{'>', '<', '+', '-', '.', ',', '[', ']'})
}

func TestTokenize(t *testing.T) {
	quickcheck.Check(t, "token", func() bool {
		tokens, err := lexical.Tokenize(strings.NewReader(quickcheck.String()))
		if err != nil {
			return false
		}

		for _, token := range tokens {
			if token < lexical.DecDataPointerToken || token > lexical.WhileCloseToken {
				return false
			}
		}
		return true
	})

	quickcheck.Check(t, "len", func() bool {
		lexicallyValidCode := GenBrainFckWithoutComments()

		tokens, err := lexical.Tokenize(strings.NewReader(lexicallyValidCode))
		if err != nil {
			return false
		}

		return len(lexicallyValidCode) == len(tokens)
	})
}
