package operations

import TYPES "gorth/types"

func Push(operand TYPES.Operand) TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OPPUSH, operand)
}

func Plus() TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OPPLUS)
}

func Dump() TYPES.InsTUPLE {
	return append(make(TYPES.InsTUPLE, 0), TYPES.OPDUMP)
}
