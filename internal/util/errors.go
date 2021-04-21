package util

func PanicError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
