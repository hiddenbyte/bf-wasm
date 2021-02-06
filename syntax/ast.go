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

// Statement a brainfck statement
type Statement interface {
	WhileStatement(pattern func(StatementList))
	SingleCharStatement(pattern func(byte))
}

// SingleCharStatement single character statement
type SingleCharStatement struct {
	StatementToken byte
}

// WhileStatement does nothing
func (sc SingleCharStatement) WhileStatement(_ func(StatementList)) {
	return
}

// SingleCharStatement single char statement case
func (sc SingleCharStatement) SingleCharStatement(pattern func(byte)) {
	pattern(sc.StatementToken)
	return
}

// WhileStatement while loop statement
type WhileStatement struct {
	statements StatementList
}

// WhileStatement while statement case
func (sc WhileStatement) WhileStatement(pattern func(StatementList)) {
	pattern(sc.statements)
	return
}

// SingleCharStatement does nothing
func (sc WhileStatement) SingleCharStatement(_ func(byte)) {
	return
}
