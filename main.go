package main

import (
	OP "gorth/operations"
	SIM "gorth/simulation"
	TYPES "gorth/types"
)

var program TYPES.Program

func init() {
	program = TYPES.Program{
		Operations: make([]TYPES.InsTUPLE, 0),
	}
}

func main() {
	program.Operations = append(program.Operations, OP.Push(34), OP.Push(35), OP.Plus(), OP.Dump())
	SIM.Simulate(program)
}
