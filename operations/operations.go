package operations

import (
	TYPES "gorth/types"
)

func Push(operand TYPES.Operand) TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OPPUSH, operand)
}

func Plus() TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OPPLUS)
}

func Dump() TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OPDUMP)
}

func StackPop(stack []TYPES.Operand) []TYPES.Operand {
	return stack[:len(stack)-1]
}

func GetLastNDrop(stack *[]TYPES.Operand) interface{} {
	cp := *stack
	ret := cp[len(cp)-1]
	cp = cp[:len(cp)-1]
	*stack = cp
	return ret
}
