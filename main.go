package main

import (
	"fmt"
	COM "gorth/compilation"
	OP "gorth/operations"
	SIM "gorth/simulation"
	"log"
	"os"
)

func usage(prog string) {
	fmt.Printf("USAGE: %s <SUBCOMMAND> [ARGS]\n", prog)
	fmt.Println("SUBCOMMANDS:")
	fmt.Println("\tsim <file>\tSimulate the program")
	fmt.Println("\tcom <file>\tCompile the program")
}

func main() {
	argv := append(make([]interface{}, 0), os.Args)

	program_name := OP.Uncons(&argv)
	if len(argv[0].([]string)) < 2 {
		usage(program_name.(string))
		os.Exit(1)
	}

	subcommand := OP.Uncons(&argv).(string) // remove subcommand
	pathToProgram := OP.Uncons(&argv)       // remove target file
	program := OP.LoadProgramFromFile(pathToProgram.(string))

	if subcommand == "sim" {
		SIM.Simulate(program)
	} else if subcommand == "com" {
		COM.Compile(program, "output.asm")
		err := COM.ToASM("output")
		if err != nil {
			log.Fatalf("Error: %v", err.Error())
		}
		log.Println("Output asm compiled and linked")
	} else {
		log.Printf("unsupported subcommand %s\n", subcommand)
	}
}
