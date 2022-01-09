package errors

type Error int

const (
	ESliceEmpty Error = iota
)

var Errors map[Error]string = map[Error]string{
	ESliceEmpty: "No more items in the stack",
}

func CheckErr(err Error) {
	if err == ESliceEmpty {
		panic(Errors[err])
	}
}
