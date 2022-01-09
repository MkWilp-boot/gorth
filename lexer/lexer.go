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
		ASSERT.Assert(TYPES.CountOps == 10, "Exhaustive handling of operations in CrossReferenceBlocks")
		switch op[0] {
		case TYPES.OpIf:
			stack = append(stack, ip)
		case TYPES.OpElse:
			ifIp, err := OP.GetLastNDrop(&stack)
			if err == ERR.ESliceEmpty {
				panic(ERR.Errors[err])
			}
			ASSERT.Assert(program.Operations[ifIp.(int)][0] == TYPES.OpIf, "Else can only be used in if blocks")

			program.Operations[ifIp.(int)] = append(make(TYPES.InsTUPLE, 0), TYPES.OpIf, ip)
			stack = append(stack, ip)
		case TYPES.OpEnd:
			blockIp, err := OP.GetLastNDrop(&stack)
			if err == ERR.ESliceEmpty {
				panic(ERR.Errors[err])
			}
			if program.Operations[blockIp.(int)][0] == TYPES.OpIf || program.Operations[blockIp.(int)][0] == TYPES.OpElse {
				program.Operations[blockIp.(int)] = append(make(TYPES.InsTUPLE, 0), program.Operations[blockIp.(int)][0], ip)
			} else {
				ASSERT.Assert(false, "End can only close if-else blocks")
			}
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
	lines := make([]interface{}, 0)
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
