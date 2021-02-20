package text

import (
	"fmt"
	"io"
)

// SymbolicExpression s-expression
type SymbolicExpression interface {
	Print(w io.Writer) error
}

// SymbolicExpressionList list of s-expressions
type SymbolicExpressionList []SymbolicExpression

// Print prints the s-expression
func (l SymbolicExpressionList) Print(w io.Writer) error {
	if len(l) == 0 {
		_, err := io.WriteString(w, "()")
		return err
	}

	_, err := io.WriteString(w, "(")
	if err != nil {
		return err
	}

	l[0].Print(w)
	for _, expression := range l[1:] {
		_, err := io.WriteString(w, " ")
		if err != nil {
			return err
		}

		err = expression.Print(w)
		if err != nil {
			return err
		}
	}

	_, err = io.WriteString(w, ")")
	return err
}

// AtomIdentifier atom as an identifier
type AtomIdentifier string

func (s AtomIdentifier) String() string {
	return fmt.Sprintf("$%s", string(s))
}

// Print prints the s-expression
func (s AtomIdentifier) Print(w io.Writer) error {
	_, err := io.WriteString(w, s.String())
	if err != nil {
		return err
	}
	return err
}

// AtomString atom as a string literal
type AtomString string

func (s AtomString) String() string {
	return fmt.Sprintf("\"%s\"", string(s))
}

// Print prints the s-expression
func (s AtomString) Print(w io.Writer) error {
	_, err := io.WriteString(w, s.String())
	if err != nil {
		return err
	}
	return err
}

// Atom simple atom
type Atom string

// Print prints the s-expression
func (s Atom) Print(w io.Writer) error {
	_, err := io.WriteString(w, string(s))
	if err != nil {
		return err
	}
	return err
}

// Value types

// I32 signed int32 value type
const I32 Atom = "i32"
