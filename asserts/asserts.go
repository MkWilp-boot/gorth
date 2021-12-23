package asserts

type assertion bool

func Assert(condition assertion, message string) {
	if message == "" {
		message = "Got false assetion"
	}

	if !condition {
		panic(message)
	}
}
