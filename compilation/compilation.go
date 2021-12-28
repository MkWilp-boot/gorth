package compilation

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	ASSERT "gorth/asserts"
	TYPES "gorth/types"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ToASM(file string) error {
	cmd := exec.Command("nasm", "-felf64", file+".asm")

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err := cmd.Run()
	if err != nil {
		return errors.New(errb.String())
	}
	log.Println(outb.String())

	cmd = exec.Command("ld", "-o", file, file+".o")

	cmd.Stdout = &outb
	cmd.Stderr = &errb

	err = cmd.Run()
	if err != nil {
		return errors.New(errb.String())
	}
	log.Println(outb.String())
	return nil
}

func Compile(program TYPES.Program, outfilePath string) {
	output, err := os.Create(outfilePath)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(output)
	writer.WriteString("BITS 64\n")
	writer.WriteString("section .text\n")
	writer.WriteString("dump:\n")
	writer.WriteString("\tmov     r9, -3689348814741910323\n")
	writer.WriteString("\tsub     rsp, 40\n")
	writer.WriteString("\tmov     BYTE [rsp+31], 10\n")
	writer.WriteString("\tlea     rcx, [rsp+30]\n")
	writer.WriteString(".L2:\n")
	writer.WriteString("\tmov     rax, rdi\n")
	writer.WriteString("\tlea     r8, [rsp+32]\n")
	writer.WriteString("\tmul     r9\n")
	writer.WriteString("\tmov     rax, rdi\n")
	writer.WriteString("\tsub     r8, rcx\n")
	writer.WriteString("\tshr     rdx, 3\n")
	writer.WriteString("\tlea     rsi, [rdx+rdx*4]\n")
	writer.WriteString("\tadd     rsi, rsi\n")
	writer.WriteString("\tsub     rax, rsi\n")
	writer.WriteString("\tadd     eax, 48\n")
	writer.WriteString("\tmov     BYTE [rcx], al\n")
	writer.WriteString("\tmov     rax, rdi\n")
	writer.WriteString("\tmov     rdi, rdx\n")
	writer.WriteString("\tmov     rdx, rcx\n")
	writer.WriteString("\tsub     rcx, 1\n")
	writer.WriteString("\tcmp     rax, 9\n")
	writer.WriteString("\tja      .L2\n")
	writer.WriteString("\tlea     rax, [rsp+32]\n")
	writer.WriteString("\tmov     edi, 1\n")
	writer.WriteString("\tsub     rdx, rax\n")
	writer.WriteString("\txor     eax, eax\n")
	writer.WriteString("\tlea     rsi, [rsp+32+rdx]\n")
	writer.WriteString("\tmov     rdx, r8\n")
	writer.WriteString("\tmov	  rax, 1\n")
	writer.WriteString("\tsyscall\n")
	writer.WriteString("\tadd     rsp, 40\n")
	writer.WriteString("\tret\n")

	writer.WriteString("global _start\n")
	writer.WriteString("_start:\n")

	for ip := 0; ip < len(program.Operations); ip++ {
		op := program.Operations[ip]
		ASSERT.Assert(TYPES.CountOps == 8, "Exhaustive handling of operations in simulation")
		switch op[0] {
		case TYPES.OpPush:
			writer.WriteString(fmt.Sprintf("\t; push %d\n", op[1]))
			writer.WriteString(fmt.Sprintf("\tpush     %d\n", op[1]))
		case TYPES.OpPlus:
			writer.WriteString("\t; plus \n")
			writer.WriteString("\tpop     rax\n")
			writer.WriteString("\tpop     rbx\n")
			writer.WriteString("\tadd     rax, rbx\n")
			writer.WriteString("\tpush    rax\n")
		case TYPES.OpMinus:
			writer.WriteString("\t; minus \n")
			writer.WriteString("\tpop     rax\n")
			writer.WriteString("\tpop     rbx\n")
			writer.WriteString("\tsub     rbx, rax\n")
			writer.WriteString("\tpush    rbx\n")
		case TYPES.OpDump:
			writer.WriteString("\t; dump\n")
			writer.WriteString("\tpop     rdi\n")
			writer.WriteString("\tcall    dump\n")
		case TYPES.OpEqual:
			writer.WriteString("\t; equal\n")
			writer.WriteString("\tmov rcx, 0\n")
			writer.WriteString("\tmov rdx, 1\n")
			writer.WriteString("\tpop rax\n")
			writer.WriteString("\tpop rbx\n")
			writer.WriteString("\tcmp rax, rbx\n")
			writer.WriteString("\tcmove rcx, rdx\n")
			writer.WriteString("\tpush rcx\n")
		case TYPES.OpIf:
			writer.WriteString(fmt.Sprintf("\t; If of %d\n", op[1]))
			writer.WriteString("\tpop rax\n")
			writer.WriteString("\ttest rax, rax\n")

			ASSERT.Assert(len(op) >= 2, "During compilation, If instruction does not have a reference to an End block")

			writer.WriteString(fmt.Sprintf("\tjz addr_%d\n", op[1]))
		case TYPES.OpElse:
			writer.WriteString("\t; Else\n")
			ASSERT.Assert(len(op) >= 2, "During compilation, Else instruction does not have a reference to an If block")

			writer.WriteString(fmt.Sprintf("\tjmp addr_%d\n", op[1]))
			writer.WriteString(fmt.Sprintf("addr_%d:\n", ip))
		case TYPES.OpEnd:
			writer.WriteString(fmt.Sprintf("\t; End of %d\n", ip))
			writer.WriteString(fmt.Sprintf("addr_%d:\n", ip))
		default:
			ASSERT.Assert(false, "unreachable")
		}
	}

	writer.WriteString("\tmov     rax, 60\n")
	writer.WriteString("\tmov     rdi, 0\n")
	writer.WriteString("\tsyscall\n")
	writer.Flush()
	output.Close()

	err = ToASM(strings.Split(outfilePath, ".")[0])
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}
	log.Println("Output asm compiled and linked")
}
