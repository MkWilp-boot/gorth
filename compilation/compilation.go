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
	writer.WriteString("    mov     r9, -3689348814741910323\n")
	writer.WriteString("    sub     rsp, 40\n")
	writer.WriteString("    mov     BYTE [rsp+31], 10\n")
	writer.WriteString("    lea     rcx, [rsp+30]\n")
	writer.WriteString(".L2:\n")
	writer.WriteString("    mov     rax, rdi\n")
	writer.WriteString("    lea     r8, [rsp+32]\n")
	writer.WriteString("    mul     r9\n")
	writer.WriteString("    mov     rax, rdi\n")
	writer.WriteString("    sub     r8, rcx\n")
	writer.WriteString("    shr     rdx, 3\n")
	writer.WriteString("    lea     rsi, [rdx+rdx*4]\n")
	writer.WriteString("    add     rsi, rsi\n")
	writer.WriteString("    sub     rax, rsi\n")
	writer.WriteString("    add     eax, 48\n")
	writer.WriteString("    mov     BYTE [rcx], al\n")
	writer.WriteString("    mov     rax, rdi\n")
	writer.WriteString("    mov     rdi, rdx\n")
	writer.WriteString("    mov     rdx, rcx\n")
	writer.WriteString("    sub     rcx, 1\n")
	writer.WriteString("    cmp     rax, 9\n")
	writer.WriteString("    ja      .L2\n")
	writer.WriteString("    lea     rax, [rsp+32]\n")
	writer.WriteString("    mov     edi, 1\n")
	writer.WriteString("    sub     rdx, rax\n")
	writer.WriteString("    xor     eax, eax\n")
	writer.WriteString("    lea     rsi, [rsp+32+rdx]\n")
	writer.WriteString("    mov     rdx, r8\n")
	writer.WriteString("    mov	    rax, 1\n")
	writer.WriteString("    syscall\n")
	writer.WriteString("    add     rsp, 40\n")
	writer.WriteString("    ret\n")

	writer.WriteString("global _start\n")
	writer.WriteString("_start:\n")

	for ip := 0; ip < len(program.Operations); ip++ {
		op := program.Operations[ip]
		ASSERT.Assert(TYPES.CountOps == 12, "Exhaustive handling of operations in compilation")
		writer.WriteString(fmt.Sprintf("addr_%d:\n", ip))
		switch op[0] {
		case TYPES.OpPush:
			writer.WriteString(fmt.Sprintf("    ; push %d\n", op[1]))
			writer.WriteString(fmt.Sprintf("    push %d\n", op[1]))
		case TYPES.OpPlus:
			writer.WriteString("    ; plus \n")
			writer.WriteString("    pop     rax\n")
			writer.WriteString("    pop     rbx\n")
			writer.WriteString("    add     rax, rbx\n")
			writer.WriteString("    push    rax\n")
		case TYPES.OpMinus:
			writer.WriteString("    ; minus \n")
			writer.WriteString("    pop     rax\n")
			writer.WriteString("    pop     rbx\n")
			writer.WriteString("    sub     rbx, rax\n")
			writer.WriteString("    push    rbx\n")
		case TYPES.OpDump:
			writer.WriteString("    ; dump\n")
			writer.WriteString("    pop     rdi\n")
			writer.WriteString("    call    dump\n")
		case TYPES.OpEqual:
			writer.WriteString("    ; equal\n")
			writer.WriteString("    mov rcx, 0\n")
			writer.WriteString("    mov rdx, 1\n")
			writer.WriteString("    pop rax\n")
			writer.WriteString("    pop rbx\n")
			writer.WriteString("    cmp rax, rbx\n")
			writer.WriteString("    cmove rcx, rdx\n")
			writer.WriteString("    push rcx\n")
		case TYPES.OpGT:
			writer.WriteString("    ; GT\n")
			writer.WriteString("    mov rcx, 0\n")
			writer.WriteString("    mov rdx, 1\n")
			writer.WriteString("    pop rbx\n")
			writer.WriteString("    pop rax\n")
			writer.WriteString("    cmp rax, rbx\n")
			writer.WriteString("    cmovg rcx, rdx\n")
			writer.WriteString("    push rcx\n")
		case TYPES.OpWhile:
			writer.WriteString("    ; while\n")
		case TYPES.OpDo:
			writer.WriteString("    ; Do\n")
			writer.WriteString("    pop rax\n")
			writer.WriteString("    test rax, rax\n")
			ASSERT.Assert(len(op) >= 2, "During compilation, 'Do' instruction does not have a reference to an End block")

			writer.WriteString(fmt.Sprintf("    jz addr_%d\n", op[1]))
		case TYPES.OpIf:
			writer.WriteString(fmt.Sprintf("    ; If of %d\n", op[1]))
			writer.WriteString("    pop rax\n")
			writer.WriteString("    test rax, rax\n")

			ASSERT.Assert(len(op) >= 2, "During compilation, If instruction does not have a reference to an End block")

			writer.WriteString(fmt.Sprintf("    jz addr_%d\n", op[1]))
		case TYPES.OpElse:
			writer.WriteString("    ; Else\n")
			ASSERT.Assert(len(op) >= 2, "During compilation, Else instruction does not have a reference to an If block")

			writer.WriteString(fmt.Sprintf("    jmp addr_%d\n", op[1]))
		case TYPES.OpDup:
			writer.WriteString("    ; Dup\n")
			writer.WriteString("    pop rax\n")
			writer.WriteString("    push rax\n")
			writer.WriteString("    push rax\n")
		case TYPES.OpEnd:
			writer.WriteString("    ; End\n")
			ASSERT.Assert(len(op) >= 2, "During compilation, End instruction does not have a reference to an If block")
			if ip+1 != op[1] {
				writer.WriteString(fmt.Sprintf("    jmp addr_%d\n", op[1]))
			}
		default:
			ASSERT.Assert(false, "unreachable")
		}
	}

	writer.WriteString("    mov     rax, 60\n")
	writer.WriteString("    mov     rdi, 0\n")
	writer.WriteString("    syscall\n")
	writer.Flush()
	output.Close()

	err = ToASM(strings.Split(outfilePath, ".")[0])
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}
	log.Println("Output asm compiled and linked")
}
