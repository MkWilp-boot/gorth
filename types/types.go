package types

type Operation int
type Operand interface{}
type InsTUPLE []interface{}

type Program struct {
	Operations []InsTUPLE
}

type Enumerator struct {
	Index uint
	Slice interface{}
}

const (
	OpPush Operation = iota
	OpPlus
	OpMinus
	OpEqual
	OpIf
	OpElse
	OpEnd
	OpDump
	CountOps
)
