package asserts

type assertion bool

func AssertThat(condition assertion, message string) {
	if message == "" {
		message = "Got false assertion"
	}

	if !condition {
		panic(message)
	}
}
