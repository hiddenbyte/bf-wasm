package wasm

import (
	"io"
	"os"

	"github.com/hiddenbyte/bf-wasm/lexical"
	"github.com/hiddenbyte/bf-wasm/syntax"
	"github.com/hiddenbyte/bf-wasm/wasm/text"
)

// localDataPointerID data pointer identifier
const localDataPointerID = text.AtomIdentifier("pointer")

// readInputFuncID read input func identifier
const readInputFuncID = text.AtomIdentifier("readInput")

// writeOutputID read input func identifier
const writeOutputID = text.AtomIdentifier("writeOutput")

// EmitText emits WebAssembly text
func EmitText(program syntax.StatementList, w io.Writer) error {

	instructions := text.SymbolicExpressionList{
		//Prologue
		text.Int32Const(0),
		text.LocalSet(localDataPointerID),
	}

	for _, statement := range program {
		statement.SingleCharStatement(func(token byte) {
			instructions = append(instructions, ToWASMInstructions(token)...)
		})
		statement.WhileStatement(func(l syntax.StatementList) {
			// TODO: Implement
		})
	}

	return text.Module(
		// Types
		text.SymbolicExpressionList{
			text.Type(
				text.AtomIdentifier("tWriteOuput"),
				text.FuncType(text.SymbolicExpressionList{text.Param("b", text.I32)}, text.SymbolicExpressionList{}),
			),
			text.Type(
				text.AtomIdentifier("tReadInput"),
				text.FuncType(text.SymbolicExpressionList{}, text.SymbolicExpressionList{text.Result(text.I32)}),
			),
			text.Type(
				text.AtomIdentifier("tEntryPoint"),
				text.FuncType(text.SymbolicExpressionList{}, text.SymbolicExpressionList{}),
			),
		},

		// Imports
		text.SymbolicExpressionList{
			text.Import("io", "writeOutput", text.ImportDescFunc(writeOutputID, text.TypeUse("tWriteOuput"))),
			text.Import("io", "readInput", text.ImportDescFunc(readInputFuncID, text.TypeUse("tReadInput"))),
		},

		// "entry point" function
		text.SymbolicExpressionList{
			text.Function("entrypoint", text.TypeUse("tEntryPoint"), text.SymbolicExpressionList{text.Local(localDataPointerID, text.I32)}, instructions),
		},

		// Export entrypoint function as 'main'
		text.SymbolicExpressionList{text.Export("main", text.ExportDescFunc("entrypoint"))},
	).Print(os.Stdout)
}

// ToWASMInstructions maps an syntax.SingleCharStatement into a sequence of WASM instruction
func ToWASMInstructions(token byte) []text.SymbolicExpression {
	switch token {
	case lexical.IncDataPointerToken:
		return []text.SymbolicExpression{
			text.LocalGet(localDataPointerID),
			text.Int32Const(1),
			text.Add(text.I32),
			text.LocalSet(localDataPointerID),
		}
	case lexical.DecDataPointerToken:
		return []text.SymbolicExpression{
			text.LocalGet(localDataPointerID),
			text.Int32Const(1),
			text.Sub(text.I32),
			text.LocalSet(localDataPointerID),
		}
	case lexical.IncDataToken:
		return []text.SymbolicExpression{
			text.LocalGet(localDataPointerID),
			text.LocalGet(localDataPointerID),
			text.Load(text.I32),
			text.Int32Const(1),
			text.Add(text.I32),
			text.Store(text.I32),
		}
	case lexical.DecDataToken:
		return []text.SymbolicExpression{
			text.LocalGet(localDataPointerID),
			text.LocalGet(localDataPointerID),
			text.Load(text.I32),
			text.Int32Const(1),
			text.Sub(text.I32),
			text.Store(text.I32),
		}
	case lexical.InputToken:
		return []text.SymbolicExpression{
			text.LocalGet(localDataPointerID),
			text.Call(readInputFuncID),
			text.Store(text.I32),
		}
	case lexical.OutputToken:
		return []text.SymbolicExpression{
			text.LocalGet(localDataPointerID),
			text.Load(text.I32),
			text.Call(writeOutputID),
		}
	default:
		panic("wasm: Unknown syntax.SingleCharStatement statement")
	}
}
