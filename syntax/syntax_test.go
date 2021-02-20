package syntax_test

import (
	"bytes"
	"testing"

	"github.com/hiddenbyte/bf-wasm/lexical"
	"github.com/hiddenbyte/bf-wasm/quickcheck"
	"github.com/hiddenbyte/bf-wasm/syntax"
)

// BrainfckCode syntatically valid brainfck code
type BrainfckCode []byte

// singleCharStatement random single char statement
func singleCharStatement() BrainfckCode {
	return BrainfckCode{
		quickcheck.OneOf(
			lexical.DecDataPointerToken,
			lexical.DecDataToken,
			lexical.IncDataPointerToken,
			lexical.IncDataToken,
			lexical.InputToken,
			lexical.OutputToken,
		).(byte),
	}
}

// whileStatement random while statement
func whileStatement() BrainfckCode {
	code := BrainfckCode{}
	code = append(code, lexical.WhileOpenToken)
	code = append(code, validBrainfckCode()...)
	code = append(code, lexical.WhileCloseToken)
	return code
}

// statement random statement
func statement() BrainfckCode {
	s := quickcheck.OneOf(singleCharStatement, whileStatement)
	return s.(func() BrainfckCode)()
}

// validBrainfckCode random syntatically valid brainfck code
func validBrainfckCode() BrainfckCode {
	code := BrainfckCode{}

	end := quickcheck.Boolean()
	for !end {
		code = append(code, statement()...)
		end = quickcheck.Boolean()
	}

	return code
}

// revertParser syntax.Parse^-1
func revertParse(programAST syntax.StatementList) []byte {
	code := []byte{}
	for _, statement := range programAST {
		statement.SingleCharStatement(func(t byte) {
			code = append(code, t)
		})
		statement.WhileStatement(func(p syntax.StatementList) {
			code = append(code, lexical.WhileOpenToken)
			code = append(code, revertParse(p)...)
			code = append(code, lexical.WhileCloseToken)
		})
	}
	return code
}

func TestParse(t *testing.T) {
	quickcheck.AddRandomGenerator(func() interface{} { return validBrainfckCode() })

	quickcheck.Check(t, "empty", func() bool {
		statements := syntax.Parse([]byte{})
		return len(statements) == 0
	})

	quickcheck.Check(t, "reverse", func(code BrainfckCode) bool {
		statements := syntax.Parse([]byte(code))
		return bytes.Equal(revertParse(statements), []byte(code))
	})
}
