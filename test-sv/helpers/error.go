package helpers

func PanicIfError(err error) {
	if err != nil {
		println(err)
		panic(err)
	}
}
