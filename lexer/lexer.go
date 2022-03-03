package lexer

import (
	"gorth/asserts"
	"gorth/errors"
	"gorth/helpers"
	"gorth/operations"
	"gorth/types"
	"io/ioutil"
	"runtime"
	"strings"
)

var newLine string

func init() {
	if runtime.GOOS == "windows" {
		newLine = "\r\n"
	} else {
		newLine = "\n"
	}
}

func CrossReferenceBlocks(program types.Program) types.Program {
	stack := make([]types.Operand, 0)

	for ip, op := range program.Operations {
		asserts.AssertThat(types.CountOps == 15, "Exhaustive handling of operations in CrossReferenceBlocks")
		switch op[0] {
		case types.OpIf:
			stack = append(stack, ip)
		case types.OpElse:
			ifIIp, err := operations.GetLastNDrop(&stack)
			ifIp := ifIIp.(int)
			errors.CheckErr(err)
			asserts.AssertThat(program.Operations[ifIp][0] == types.OpIf, "Else can only be used in if blocks")

			program.Operations[ifIp] = append(make(types.InsTUPLE, 0), types.OpIf, ip+1)
			stack = append(stack, ip)
		case types.OpEnd:
			blockIIp, err := operations.GetLastNDrop(&stack)
			blockIp := blockIIp.(int)
			errors.CheckErr(err)
			if program.Operations[blockIp][0] == types.OpIf ||
				program.Operations[blockIp][0] == types.OpElse {

				program.Operations[blockIp] = append(make(types.InsTUPLE, 0), program.Operations[blockIp][0], ip)
				program.Operations[ip] = append(make(types.InsTUPLE, 0), types.OpEnd, ip+1)

			} else if program.Operations[blockIp][0] == types.OpDo {
				asserts.AssertThat(len(program.Operations[blockIp]) >= 2, "")
				program.Operations[ip] = append(make(types.InsTUPLE, 0), types.OpEnd, program.Operations[blockIp][1])
				program.Operations[blockIp] = append(make(types.InsTUPLE, 0), types.OpDo, ip+1)
			} else {
				asserts.AssertThat(false, "End can only close 'if','else' or 'do' blocks")
			}
		case types.OpWhile:
			stack = append(stack, ip)
		case types.OpDo:
			whileIIp, err := operations.GetLastNDrop(&stack)
			whileIp := whileIIp.(int)
			errors.CheckErr(err)
			program.Operations[ip] = append(
				make(types.InsTUPLE, 0), types.OpDo, whileIp)

			stack = append(stack, ip)
		}
	}
	return program
}

func LoadProgramFromFile(filePath string) types.Program {
	enumerate := LexFile(filePath)
	program := types.Program{
		Operations: operations.ParseTokenAsOperation(enumerate, filePath),
	}
	return CrossReferenceBlocks(program)
}

func LexFile(filePath string) []types.Vec2DString {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	lines := make([]types.Vec2DString, 0)

	for lineNumber, line := range strings.Split(string(bytes), newLine) {
		if len(line) == 0 {
			continue
		}
		enumeration := make(chan types.StringEnum)

		go helpers.EnumerateLine(line, enumeration)

		for enumeratedLine := range enumeration {
			vec2d := types.Vec2DString{
				Index:   lineNumber + 1,
				Content: enumeratedLine,
			}
			lines = append(lines, vec2d)
		}
	}
	return helpers.Enumerate(lines)
}
