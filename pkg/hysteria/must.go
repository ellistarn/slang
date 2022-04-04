package hysteria

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustAssign[T any](value T, err error) {
	if err != nil {
		panic(err)
	}
}
