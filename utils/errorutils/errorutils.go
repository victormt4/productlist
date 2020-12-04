package errorutils

func PanicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
