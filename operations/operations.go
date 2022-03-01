package operations

import (
	"fmt"
	"gorth/asserts"
	ERR "gorth/errors"
	"gorth/types"
	"strconv"
)

func ParseTokenAsOperation(tokens []types.Vec2DString, filePath string) []types.InsTUPLE {
	asserts.AssertThat(types.CountOps == 12, "Exhaustive handling of operations during parser")
	ops := make([]types.InsTUPLE, 0)

	for _, value := range tokens {

		pair := value.Content

		token := pair.Slice
		col := pair.Index
		line := value.Content.Index

		numeric, isNum := strconv.Atoi(token)
		switch {
		case token == "+":
			ops = append(ops, Plus())
		case isNum == nil:
			ops = append(ops, Push(numeric))
		case token == "-":
			ops = append(ops, Minus())
		case token == "dump":
			ops = append(ops, Dump())
		case token == "=":
			ops = append(ops, Equal())
		case token == ">":
			ops = append(ops, GT())
		case token == "if":
			ops = append(ops, If())
		case token == "end":
			ops = append(ops, End())
		case token == "else":
			ops = append(ops, Else())
		case token == "while":
			ops = append(ops, While())
		case token == "do":
			ops = append(ops, Do())
		case token == "dup":
			ops = append(ops, Dup())
		default:
			asserts.AssertThat(false, fmt.Sprintf("File %q Line %d Column %d: %q is not a valid command", filePath, line, col, token))
		}
	}
	return ops
}

func GetLastNDrop(stack *[]types.Operand) (interface{}, ERR.Error) {
	cp := *stack
	if len(cp) == 0 {
		return -1, ERR.ESliceEmpty
	}
	ret := cp[len(cp)-1]
	cp = cp[:len(cp)-1]
	*stack = cp
	return ret, -1
}

func Push(operand types.Operand) types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpPush, operand)
}

func Plus() types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpPlus)
}

func Minus() types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpMinus)
}

func Dump() types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpDump)
}

func Equal() types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpEqual)
}

func GT() types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpGT)
}

func If() types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpIf)
}

func Else() types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpElse)
}

func Dup() types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpDup)
}

func While() types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpWhile)
}

func Do() types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpDo)
}

func End() types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpEnd)
}

func Mem() types.InsTUPLE {
	return append(make(types.InsTUPLE, 0), types.OpMem)
}

func StackPop(stack []types.Operand) []types.Operand {
	return stack[:len(stack)-1]
}
