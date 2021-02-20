package syntax

/*
	Brainfck language grammar

	Program => StatementList

	StatementList => StatementList Statement

	Statement => SingleCharStatement

	SingleCharStatement => '>'
	SingleCharStatement => '<'
	SingleCharStatement => '+'
	SingleCharStatement => '-'
	SingleCharStatement => '.'
	SingleCharStatement => ','

	Statement => WhileStatement

	WhileStatement => '['  StatementList ']'
*/

// StatementList statements list
type StatementList []Statement

// Statement a statement
type Statement interface {
	WhileStatement(pattern func(StatementList))
	SingleCharStatement(pattern func(byte))
}

// SingleCharStatement single character statement
type SingleCharStatement struct {
	StatementToken byte
}

// WhileStatement while statement pattern
func (sc SingleCharStatement) WhileStatement(_ func(StatementList)) {
	return
}

// SingleCharStatement single char statement pattern
func (sc SingleCharStatement) SingleCharStatement(pattern func(byte)) {
	pattern(sc.StatementToken)
	return
}

// WhileStatement while statement
type WhileStatement struct {
	statements StatementList
}

// WhileStatement while statement pattern
func (sc WhileStatement) WhileStatement(pattern func(StatementList)) {
	pattern(sc.statements)
	return
}

// SingleCharStatement single char statement pattern
func (sc WhileStatement) SingleCharStatement(_ func(byte)) {
	return
}
