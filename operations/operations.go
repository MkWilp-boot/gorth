package operations

import (
	ASSERT "gorth/asserts"
	TYPES "gorth/types"
	"io/ioutil"
	"strconv"
	"strings"
)

func Push(operand TYPES.Operand) TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OpPush, operand)
}

func Plus() TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OpPlus)
}

func Minus() TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OpMinus)
}

func Dump() TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OpDump)
}

func StackPop(stack []TYPES.Operand) []TYPES.Operand {
	return stack[:len(stack)-1]
}

func LoadProgramFromFile(filePath string) TYPES.Program {
	program := TYPES.Program{}
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	tokens := strings.ReplaceAll(string(bytes), "\n", " ")
	content := strings.Split(tokens, " ")

	program.Operations = parseWordAsOperation(content)
	return program
}

func parseWordAsOperation(words []string) []TYPES.InsTUPLE {
	ASSERT.Assert(TYPES.CountOps == 4, "Exhaustive handling of operations in simulation")
	ops := make([]TYPES.InsTUPLE, 0)
	for _, word := range words {
		numeric, isNum := strconv.Atoi(word)

		switch {
		case word == "+":
			ops = append(ops, Plus())
		case isNum == nil:
			ops = append(ops, Push(numeric))
		case word == "-":
			ops = append(ops, Minus())
		case word == ".":
			ops = append(ops, Dump())
		}
	}
	return ops
}

func GetLastNDrop(stack *[]TYPES.Operand) interface{} {
	cp := *stack
	ret := cp[len(cp)-1]
	cp = cp[:len(cp)-1]
	*stack = cp
	return ret
}

func Uncons(slice *[]interface{}) (ret interface{}) {
	cp := *slice
	ret = cp[0].([]string)[0]
	cp[0] = cp[0].([]string)[1:]
	*slice = cp
	return
}
