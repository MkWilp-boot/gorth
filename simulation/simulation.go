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

	ip := 0
	for ip < len(program.Operations) {
		op := program.Operations[ip]

		ASSERT.Assert(TYPES.CountOps == 12, "Exhaustive handling of operations in simulation")

		switch op[0] {
		case TYPES.OpPush:
			stack = append(stack, op[1])
			ip++
		case TYPES.OpPlus:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			b, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			stack = append(stack, a.(int)+b.(int))
			ip++
		case TYPES.OpMinus:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			b, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			stack = append(stack, b.(int)-a.(int))
			ip++
		case TYPES.OpDump:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			fmt.Println(a)
			ip++
		case TYPES.OpEqual:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			b, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			s := 0
			if a == b {
				s = 1
			}
			stack = append(stack, s)
			ip++
		case TYPES.OpGT:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			b, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			s := 0
			if b.(int) > a.(int) {
				s = 1
			}
			stack = append(stack, s)
			ip++
		case TYPES.OpIf:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			if a == 0 {
				ASSERT.Assert(len(op) >= 2, "'If' instruction does not have an end block")
				ip = op[1].(int)
			} else {
				ip++
			}
		case TYPES.OpElse:
			ASSERT.Assert(len(op) >= 2, "'Else' instruction does not have an If reference block")
			ip = op[1].(int)
		case TYPES.OpDup:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			stack = append(stack, a, a)
			ip++
		case TYPES.OpWhile:
			ip++
		case TYPES.OpDo:
			a, err := OP.GetLastNDrop(&stack)
			ERR.CheckErr(err)
			if a.(int) == 0 {
				ASSERT.Assert(len(op) >= 2, "'End' instruction does not have an reference to any block")
				ip = op[1].(int)
			} else {
				ip++
			}
		case TYPES.OpEnd:
			ASSERT.Assert(len(op) >= 2, "'End' instruction does not have an reference to any block")
			ip = op[1].(int)
		default:
			ASSERT.Assert(false, "unreachable")
		}
	}
}
