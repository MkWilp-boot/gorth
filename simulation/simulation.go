package simulation

import (
	"fmt"
	ASSERT "gorth/asserts"
	OP "gorth/operations"
	TYPES "gorth/types"
)

func Simulate(program TYPES.Program) {
	stack := make([]TYPES.Operand, 0)

	for _, op := range program.Operations {
		ASSERT.Assert(TYPES.COUNTOPS == 3, "Exhaustive handling of operations in simulation")
		switch op[0] {
		case TYPES.OPPUSH:
			stack = append(stack, op[1])
		case TYPES.OPPLUS:
			a := OP.GetLastNDrop(&stack).(int)
			b := OP.GetLastNDrop(&stack).(int)
			stack = append(stack, a+b)
		case TYPES.OPDUMP:
			a := OP.GetLastNDrop(&stack)
			fmt.Println(a)
		default:
			ASSERT.Assert(false, "unreachable")
		}
	}
}
