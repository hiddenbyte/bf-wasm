package text

import "fmt"

// Param creates a func parameter
func Param(id AtomIdentifier, valType Atom) SymbolicExpression {
	return SymbolicExpressionList{
		Atom("param"),
		id,
		valType,
	}
}

// Result creates a func result
func Result(valType Atom) SymbolicExpression {
	return SymbolicExpressionList{
		Atom("result"),
		valType,
	}
}

// FuncType creates a func type
func FuncType(params []SymbolicExpression, results []SymbolicExpression) SymbolicExpression {
	expressions := make(SymbolicExpressionList, 0, 1+len(params)+len(results))
	expressions = append(expressions, Atom("func"))
	expressions = append(expressions, params...)
	expressions = append(expressions, results...)
	return expressions
}

// Type creates a type definition
func Type(id AtomIdentifier, funcType SymbolicExpression) SymbolicExpression {
	return SymbolicExpressionList{
		Atom("type"),
		id,
		funcType,
	}
}

// TypeUse creates a type use
func TypeUse(id AtomIdentifier) SymbolicExpression {
	return SymbolicExpressionList{
		Atom("type"),
		id,
	}
}

// ImportDescFunc creates an 'import description'
func ImportDescFunc(id AtomIdentifier, typeUse SymbolicExpression) SymbolicExpression {
	return SymbolicExpressionList{
		Atom("func"),
		id,
		typeUse,
	}
}

// Import creates an import
func Import(modName AtomString, name AtomString, desc SymbolicExpression) SymbolicExpression {
	return SymbolicExpressionList{
		Atom("import"),
		modName,
		name,
		desc,
	}
}

// Int32Const creates a i32.const instruction
func Int32Const(literal int32) SymbolicExpression {
	return Atom(fmt.Sprintf("i32.const %v", literal))
}

// Call creates a call instruction
func Call(funcID AtomIdentifier) SymbolicExpression {
	return Atom(fmt.Sprintf("call %s", funcID))
}

// Load creates a load instruction
func Load(valType Atom) SymbolicExpression {
	return Atom(fmt.Sprintf("%s.load", valType))
}

// Store creates a store instruction
func Store(valType Atom) SymbolicExpression {
	return Atom(fmt.Sprintf("%s.store", valType))
}

// Function creates a function definition
func Function(id AtomIdentifier, typeUse SymbolicExpression, locals []SymbolicExpression, instructions []SymbolicExpression) SymbolicExpression {
	expressions := make(SymbolicExpressionList, 0, 3+len(locals)+len(instructions))
	expressions = append(expressions, Atom("func"))
	expressions = append(expressions, id)
	expressions = append(expressions, typeUse)
	expressions = append(expressions, locals...)
	expressions = append(expressions, instructions...)
	return expressions
}

// ExportDescFunc creates an 'export description'
func ExportDescFunc(id AtomIdentifier) SymbolicExpression {
	return SymbolicExpressionList{
		Atom("func"),
		id,
	}
}

// Export creates an export
func Export(name AtomString, exportDesc SymbolicExpression) SymbolicExpression {
	return SymbolicExpressionList{
		Atom("export"),
		name,
		exportDesc,
	}
}

// Module creates a new module definition
func Module(types []SymbolicExpression, imports []SymbolicExpression, functions []SymbolicExpression, exports []SymbolicExpression) SymbolicExpression {
	expressions := make(SymbolicExpressionList, 0, 1+len(types)+len(imports)+len(functions)+len(exports))
	expressions = append(expressions, Atom("module"))
	expressions = append(expressions, types...)
	expressions = append(expressions, imports...)
	expressions = append(expressions, functions...)
	expressions = append(expressions, exports...)
	return expressions
}
