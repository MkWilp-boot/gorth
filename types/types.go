package types

type Operation int
type Operand interface{}
type InsTUPLE []interface{}

type Program struct {
	Operations []InsTUPLE
}

type StringEnum struct {
	Index int
	Slice string
}

type Vec2DString struct {
	Index   int
	Content StringEnum
}

const (
	OpPush Operation = iota
	OpPlus
	OpMinus
	OpEqual
	OpGT
	OpIf
	OpElse
	OpWhile
	OpDo
	OpDup
	OpEnd
	OpDump
	OpMem
	OpLoad
	OpStore
	CountOps
)
