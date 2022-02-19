package main

import (
	"flag"
	"fmt"
	"gorth/lexer"
	"gorth/simulation"
	"log"
	"os"
)

var help = flag.Bool("help", false, "Displays flags and gorth CLI usage")

func usage() {
	fmt.Println("USAGE: gorth <SUBCOMMAND> [ARGS]")
	fmt.Println("SUBCOMMAND:")
	fmt.Println("\tsim <file>\tSimulate the program")
}

func init() {
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		usage()
	}
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	subcommand := os.Args[1]
	pathToProgram := os.Args[2]
	program := lexer.LoadProgramFromFile(pathToProgram)

	switch subcommand {
	case "sim":
		simulation.Simulate(program)
	default:
		log.Printf("unsupported subcommand %s\n", subcommand)
	}
}
