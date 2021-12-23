package main

import (
	"fmt"
	ASSERT "gorth/asserts"
	OP "gorth/operations"
	SIM "gorth/simulation"
	TYPES "gorth/types"
	"os"
)

var program TYPES.Program

func init() {
	program = TYPES.Program{
		Operations: make([]TYPES.InsTUPLE, 0),
	}
}

func main() {
	// TODO: unhardcode program
	program.Operations = append(program.Operations,
		OP.Push(34),
		OP.Push(35),
		OP.Plus(),
		OP.Dump(),
		OP.Push(500),
		OP.Push(80),
		OP.Minus(),
		OP.Dump())

	if len(os.Args) < 2 {
		fmt.Println("USAGE: gorth <SUBCOMMAND> [ARGS]")
		fmt.Println("SUBCOMMANDS:")
		fmt.Println("\tsim\tSimulate the program")
		fmt.Println("\tcom\tCompile the program")
		os.Exit(1)
	}

	subcommand := os.Args[1]
	if subcommand == "sim" {
		SIM.Simulate(program)
	} else if subcommand == "com" {
		ASSERT.Assert(false, "Not implemented")
	}
}
