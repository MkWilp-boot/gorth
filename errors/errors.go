package errors

type Error int

const (
	ESliceEmpty Error = iota
)

var Errors map[Error]string = map[Error]string{
	ESliceEmpty: "No more items in the stack",
}
