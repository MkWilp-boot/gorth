package simulation

import (
	"fmt"
	ASSERT "gorth/asserts"
	ERR "gorth/errors"
	OP "gorth/operations"
	TYPES "gorth/types"
)

func Simulate(program TYPES.Program) {
	stack := make([]TYPES.Operand, 0)

	for ip := 0; ip < len(program.Operations); ip++ {
		op := program.Operations[ip]

		ASSERT.Assert(TYPES.CountOps == 9, "Exhaustive handling of operations in simulation")

		switch op[0] {
		case TYPES.OpPush:
			stack = append(stack, op[1])
		case TYPES.OpPlus:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			b, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			stack = append(stack, a.(int)+b.(int))
		case TYPES.OpMinus:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			b, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			stack = append(stack, b.(int)-a.(int))
		case TYPES.OpDump:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			fmt.Println(a)
		case TYPES.OpEqual:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			b, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			var s int = 0
			if a == b {
				s = 1
			}
			stack = append(stack, s)
		case TYPES.OpIf:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			if a == 0 {
				ASSERT.Assert(len(op) >= 2, "'If' instruction does not have an end block")
				ip = op[1].(int)
			}
		case TYPES.OpElse:
			ASSERT.Assert(len(op) >= 2, "'Else' instruction does not have an If reference block")
			ip = op[1].(int)
		case TYPES.OpDup:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			stack = append(stack, a, a)
		case TYPES.OpEnd:
		default:
			ASSERT.Assert(false, "unreachable")
		}
	}
}
