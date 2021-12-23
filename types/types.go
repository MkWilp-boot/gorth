package types

type Operation int
type Operand interface{}
type InsTUPLE []interface{}

type Program struct {
	Operations []InsTUPLE
}

const (
	OpPush Operation = iota
	OpPlus
	OpMinus
	OpDump
	CountOps
)
