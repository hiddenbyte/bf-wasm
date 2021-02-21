package wasm

import (
	"fmt"
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
	instructions = append(instructions, statementListToWASMInstructions(program, 0)...)

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

// statementListToWASMInstructions maps an syntax.StatementList into a sequence of WASM instruction
func statementListToWASMInstructions(p syntax.StatementList, depth int) []text.SymbolicExpression {
	instructions := text.SymbolicExpressionList{}
	for _, statement := range p {
		statement.SingleCharStatement(func(token byte) {
			instructions = append(instructions, toWASMInstructions(token)...)
		})

		statement.WhileStatement(func(l syntax.StatementList) {
			loopLabel := text.AtomIdentifier(fmt.Sprintf("L%v", depth))
			whileInstructions := statementListToWASMInstructions(l, depth+1)
			whileInstructions = append(whileInstructions, text.LocalGet(localDataPointerID))
			whileInstructions = append(whileInstructions, text.Load(text.I32))
			whileInstructions = append(whileInstructions, text.BranchIf(loopLabel))
			instructions = append(instructions, text.LocalGet(localDataPointerID))
			instructions = append(instructions, text.Load(text.I32))
			instructions = append(instructions, text.IfBlock(text.LoopBlock(loopLabel, whileInstructions), []text.SymbolicExpression{})...)
		})
	}
	return instructions
}

// toWASMInstructions maps an syntax.SingleCharStatement into a sequence of WASM instruction
func toWASMInstructions(token byte) []text.SymbolicExpression {
	switch token {
	case lexical.IncDataPointerToken:
		return []text.SymbolicExpression{
			text.LocalGet(localDataPointerID),
			text.Int32Const(4),
			text.Add(text.I32),
			text.LocalSet(localDataPointerID),
		}
	case lexical.DecDataPointerToken:
		return []text.SymbolicExpression{
			text.LocalGet(localDataPointerID),
			text.Int32Const(4),
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
