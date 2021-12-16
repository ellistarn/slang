package do

func OrDie(f func() error) {
	if err := f(); err != nil {
		panic(err)
	}
}
