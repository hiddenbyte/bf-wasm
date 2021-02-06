package syntax

import "github.com/hiddenbyte/bf-wasm/lexical"

type tokenStream struct {
	tokens []byte
	curr   int
}

func (ts tokenStream) peek() byte {
	return ts.tokens[ts.curr]
}

func (ts tokenStream) consume() tokenStream {
	return tokenStream{ts.tokens, ts.curr + 1}
}

func (ts tokenStream) isEmpty() bool {
	return ts.curr >= len(ts.tokens)
}

// Parse tranforms a slice of brainfck code tokens into brainfck ast
func Parse(tokens []byte) StatementList {
	statements, _ := newStatementList(tokenStream{tokens: tokens})
	return statements
}

// newStatementList ...
func newStatementList(tokens tokenStream) (StatementList, tokenStream) {
	statements := StatementList{}

	for !tokens.isEmpty() {
		token := tokens.peek()
		switch token {
		case lexical.WhileOpenToken:
			statement, consumedTokens := newWhileStatement(tokens)
			tokens = consumedTokens
			statements = append(statements, statement)
		case lexical.WhileCloseToken:
			return statements, tokens.consume()
		default:
			statement, consumedTokens := newSingleCharStatement(tokens)
			tokens = consumedTokens
			statements = append(statements, statement)
		}
	}

	return statements, tokens
}

// newSingleCharStatement ...
func newSingleCharStatement(tokens tokenStream) (SingleCharStatement, tokenStream) {
	return SingleCharStatement{tokens.peek()}, tokens.consume()
}

// newWhileStatement ...
func newWhileStatement(tokens tokenStream) (WhileStatement, tokenStream) {
	statements, tokens := newStatementList(tokens.consume())
	return WhileStatement{statements}, tokens
}
