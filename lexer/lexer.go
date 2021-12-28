package lexer

import (
	ASSERT "gorth/asserts"
	ERR "gorth/errors"
	HELP "gorth/helpers"
	OP "gorth/operations"
	TYPES "gorth/types"
	"io/ioutil"
	"strings"
)

func CrossReferenceBlocks(program TYPES.Program) TYPES.Program {
	stack := make([]TYPES.Operand, 0)

	for ip, op := range program.Operations {
		ASSERT.Assert(TYPES.CountOps == 7, "Exhaustive handling of operations in CrossReferenceProgram")
		switch op[0] {
		case TYPES.OpIf:
			stack = append(stack, ip)
		case TYPES.OpEnd:
			ifIp, err := OP.GetLastNDrop(&stack)
			if err == ERR.ESliceEmpty {
				panic(ERR.Errors[err])
			}
			ASSERT.Assert(program.Operations[ifIp.(int)][0] == TYPES.OpIf, "End can only close if blocks")
			program.Operations[ifIp.(int)] = append(make(TYPES.InsTUPLE, 0), TYPES.OpIf, ip)
		}
	}
	return program
}

func LoadProgramFromFile(filePath string) TYPES.Program {
	enumerate := LexFile(filePath)
	program := TYPES.Program{
		Operations: OP.ParseTokenAsOperation(enumerate, filePath),
	}
	return CrossReferenceBlocks(program)
}

func LexFile(filePath string) []TYPES.Enumerator {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	lines := make(TYPES.Enumerable, 0)
	for lineNumber, line := range strings.Split(string(bytes), "\n") {
		if len(line) != 0 {
			enumeration := make(chan TYPES.Enumerator)
			go HELP.EnumerateLine(line, enumeration)
			for enumeratedLine := range enumeration {
				lines = append(lines, TYPES.Enumerator{
					Index: uint(lineNumber + 1),
					Slice: enumeratedLine,
				})
			}
		}
	}
	return HELP.Enumerate(lines)
}
