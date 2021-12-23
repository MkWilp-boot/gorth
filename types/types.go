package types

type Operation int
type Operand interface{}
type InsTUPLE []interface{}

type Program struct {
	Operations []InsTUPLE
}

const (
	OPPUSH Operation = iota
	OPPLUS
	OPDUMP
	COUNTOPS
)
