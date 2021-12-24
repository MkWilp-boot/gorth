package compilation

import (
	"bufio"
	"fmt"
	ASSERT "gorth/asserts"
	TYPES "gorth/types"
	"log"
	"os"
	"os/exec"
)

func ToASM(file string) error {
	out, err := exec.Command("nasm", "-felf64", file+".asm").Output()
	if err != nil {
		return err
	}
	log.Println(string(out))

	out, err = exec.Command("ld", "-o", file, file+".o").Output()
	if err != nil {
		return err
	}
	log.Println(string(out))

	return nil
}

func Compile(program TYPES.Program, outfilePath string) {
	output, err := os.Create(outfilePath)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(output)
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

	for _, op := range program.Operations {
		ASSERT.Assert(TYPES.CountOps == 4, "Exhaustive handling of operations in simulation")
		switch op[0] {
		case TYPES.OpPush:
			writer.WriteString(fmt.Sprintf("\t;; push %d\n", op[1]))
			writer.WriteString(fmt.Sprintf("\tpush     %d\n", op[1]))
		case TYPES.OpPlus:
			writer.WriteString("\t;; plus \n")
			writer.WriteString("\tpop     rax\n")
			writer.WriteString("\tpop     rbx\n")
			writer.WriteString("\tadd     rax, rbx\n")
			writer.WriteString("\tpush    rax\n")
		case TYPES.OpMinus:
			writer.WriteString("\t;; minus \n")
			writer.WriteString("\tpop     rax\n")
			writer.WriteString("\tpop     rbx\n")
			writer.WriteString("\tsub     rbx, rax\n")
			writer.WriteString("\tpush    rbx\n")
		case TYPES.OpDump:
			writer.WriteString("\t;; dump %d\n")
			writer.WriteString("\tpop     rdi\n")
			writer.WriteString("\tcall    dump\n")
		default:
			ASSERT.Assert(false, "unreachable")
		}
	}

	writer.WriteString("\tmov     rax, 60\n")
	writer.WriteString("\tmov     rdi, 0\n")
	writer.WriteString("\tsyscall\n")
	writer.Flush()
	output.Close()
}
