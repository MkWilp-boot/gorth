package types

type Operation int
type Content interface{}
type Operand interface{}
type Enumerable []interface{}
type InsTUPLE []interface{}

type Program struct {
	Operations []InsTUPLE
}

type Enumerator struct {
	Index uint
	Slice Content
}

const (
	OpPush Operation = iota
	OpPlus
	OpMinus
	OpEqual
	OpIf
	OpEnd
	OpDump
	CountOps
)
