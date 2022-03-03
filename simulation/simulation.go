package simulation

import (
	"fmt"
	"gorth/asserts"
	"gorth/errors"
	"gorth/operations"
	"gorth/types"
)

const memCapacity = 640000

func Simulate(program types.Program) {
	stack := make([]types.Operand, 0)
	mem := make([]byte, memCapacity)

	ip := 0
	for ip < len(program.Operations) {
		op := program.Operations[ip]

		asserts.AssertThat(types.CountOps == 15, "Exhaustive handling of operations in simulation")

		switch op[0] {
		case types.OpPush:
			stack = append(stack, op[1])
			ip++
		case types.OpPlus:
			a, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			b, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)

			stack = append(stack, a.(int)+b.(int))
			ip++
		case types.OpMinus:
			a, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			b, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			stack = append(stack, b.(int)-a.(int))
			ip++
		case types.OpDump:
			a, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			fmt.Println(a)
			ip++
		case types.OpEqual:
			a, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			b, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			s := 0
			if a == b {
				s = 1
			}
			stack = append(stack, s)
			ip++
		case types.OpGT:
			a, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			b, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			s := 0
			if b.(int) > a.(int) {
				s = 1
			}
			stack = append(stack, s)
			ip++
		case types.OpIf:
			a, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			if a == 0 {
				asserts.AssertThat(len(op) >= 2, "'If' instruction does not have an end block")
				ip = op[1].(int)
			} else {
				ip++
			}
		case types.OpElse:
			asserts.AssertThat(len(op) >= 2, "'Else' instruction does not have an If reference block")
			ip = op[1].(int)
		case types.OpDup:
			a, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			stack = append(stack, a, a)
			ip++
		case types.OpWhile:
			ip++
		case types.OpDo:
			a, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			if a.(int) == 0 {
				asserts.AssertThat(len(op) >= 2, "'End' instruction does not have an reference to any block")
				ip = op[1].(int)
			} else {
				ip++
			}
		case types.OpEnd:
			asserts.AssertThat(len(op) >= 2, "'End' instruction does not have an reference to any block")
			ip = op[1].(int)
		case types.OpMem:
			stack = append(stack, 0)
			ip++
		case types.OpStore:
			value, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			addr, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)

			mem[addr.(int)] = uint8(value.(int))
			ip++
		case types.OpLoad:
			addr, err := operations.GetLastNDrop(&stack)
			errors.CheckErr(err)
			stack = append(stack, string(mem[addr.(int)]))
			ip++
		default:
			asserts.AssertThat(false, "unreachable")
		}
	}
}
