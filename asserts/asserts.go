package asserts

type assertion bool

func Assert(condition assertion, message string) {
	if !condition {
		panic(message)
	}
}
