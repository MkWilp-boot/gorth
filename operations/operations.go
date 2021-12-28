package operations

import (
	"fmt"
	ASSERT "gorth/asserts"
	ERR "gorth/errors"
	TYPES "gorth/types"
	"strconv"
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

func Equal() TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OpEqual)
}

func If() TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OpIf)
}

func Else() TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OpElse)
}

func End() TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OpEnd)
}

func StackPop(stack []TYPES.Operand) []TYPES.Operand {
	return stack[:len(stack)-1]
}

func ParseTokenAsOperation(tokens []TYPES.Enumerator, filePath string) []TYPES.InsTUPLE {
	ASSERT.Assert(TYPES.CountOps == 8, "Exhaustive handling of operations during parser")
	ops := make([]TYPES.InsTUPLE, 0)

	for _, value := range tokens {

		pair := value.Slice.(TYPES.Enumerator).Slice.(TYPES.Enumerator)
		token := pair.Slice.(string)
		col := pair.Index
		line := value.Slice.(TYPES.Enumerator).Index

		numeric, isNum := strconv.Atoi(token)
		switch {
		case token == "+":
			ops = append(ops, Plus())
		case isNum == nil:
			ops = append(ops, Push(numeric))
		case token == "-":
			ops = append(ops, Minus())
		case token == ".":
			ops = append(ops, Dump())
		case token == "=":
			ops = append(ops, Equal())
		case token == "if":
			ops = append(ops, If())
		case token == "end":
			ops = append(ops, End())
		case token == "else":
			ops = append(ops, Else())
		default:
			ASSERT.Assert(false, fmt.Sprintf("File %q Line %d Column %d: %q is not a valid command", filePath, line, col, token))
		}
	}
	return ops
}

func GetLastNDrop(stack *[]TYPES.Operand) (interface{}, ERR.Error) {
	cp := *stack
	if len(cp) == 0 {
		return -1, ERR.ESliceEmpty
	}
	ret := cp[len(cp)-1]
	cp = cp[:len(cp)-1]
	*stack = cp
	return ret, -1
}

func Uncons(slice *[]interface{}) (ret interface{}) {
	cp := *slice
	ret = cp[0].([]string)[0]
	cp[0] = cp[0].([]string)[1:]
	*slice = cp
	return
}
