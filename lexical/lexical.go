package lexical

import (
	"io"
)

const (
	// IncDataPointerToken 'increment data pointer' token
	IncDataPointerToken byte = iota

	// DecDataPointerToken 'decrement data pointer' token
	DecDataPointerToken

	// IncDataToken 'increment data' token
	IncDataToken

	// DecDataToken 'decrement data' token
	DecDataToken

	// OutputToken 'output current pointer data' token
	OutputToken

	// InputToken 'read input data' token
	InputToken

	// WhileOpenToken 'while start' token
	WhileOpenToken

	// WhileCloseToken 'while end' token
	WhileCloseToken
)

const bfFileReadBufferSize = 1024

// Tokenize tokenize the
func Tokenize(sourceCode io.Reader) ([]byte, error) {
	readBuffer := make([]byte, bfFileReadBufferSize)

	tokens := []byte{}
	for {
		count, err := sourceCode.Read(readBuffer)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		for i := 0; i < count; i++ {
			token, valid := tokenize(readBuffer[i])
			if valid {
				tokens = append(tokens, token)
			}
		}
	}

	return tokens, nil
}

// tokenize converts a char into a token
func tokenize(c byte) (byte, bool) {
	switch c {
	case '>':
		return IncDataPointerToken, true
	case '<':
		return DecDataPointerToken, true
	case '+':
		return IncDataToken, true
	case '-':
		return DecDataToken, true
	case '.':
		return OutputToken, true
	case ',':
		return InputToken, true
	case '[':
		return WhileOpenToken, true
	case ']':
		return WhileCloseToken, true
	default:
		return 0, false
	}
}
