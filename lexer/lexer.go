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
		ASSERT.Assert(TYPES.CountOps == 12, "Exhaustive handling of operations in CrossReferenceBlocks")
		switch op[0] {
		case TYPES.OpIf:
			stack = append(stack, ip)
		case TYPES.OpElse:
			ifIIp, err := OP.GetLastNDrop(&stack)
			ifIp := ifIIp.(int)
			ERR.CheckErr(err)
			ASSERT.Assert(program.Operations[ifIp][0] == TYPES.OpIf, "Else can only be used in if blocks")

			program.Operations[ifIp] = append(make(TYPES.InsTUPLE, 0), TYPES.OpIf, ip+1)
			stack = append(stack, ip)
		case TYPES.OpEnd:
			blockIIp, err := OP.GetLastNDrop(&stack)
			blockIp := blockIIp.(int)
			ERR.CheckErr(err)
			if program.Operations[blockIp][0] == TYPES.OpIf ||
				program.Operations[blockIp][0] == TYPES.OpElse {

				program.Operations[blockIp] = append(make(TYPES.InsTUPLE, 0), program.Operations[blockIp][0], ip)
				program.Operations[ip] = append(make(TYPES.InsTUPLE, 0), TYPES.OpEnd, ip+1)

			} else if program.Operations[blockIp][0] == TYPES.OpDo {
				ASSERT.Assert(len(program.Operations[blockIp]) >= 2, "")
				program.Operations[ip] = append(make(TYPES.InsTUPLE, 0), TYPES.OpEnd, program.Operations[blockIp][1])
				program.Operations[blockIp] = append(make(TYPES.InsTUPLE, 0), TYPES.OpDo, ip+1)
			} else {
				ASSERT.Assert(false, "End can only close 'if','else' or 'do' blocks")
			}
		case TYPES.OpWhile:
			stack = append(stack, ip)
		case TYPES.OpDo:
			whileIIp, err := OP.GetLastNDrop(&stack)
			whileIp := whileIIp.(int)
			ERR.CheckErr(err)
			program.Operations[ip] = append(
				make(TYPES.InsTUPLE, 0), TYPES.OpDo, whileIp)

			stack = append(stack, ip)
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
