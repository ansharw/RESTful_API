package helper

import "belajar-rest-api/exception"

func PanicIfError(err error) {
	if err != nil {
		panic(exception.NewInternalServerError(err))
	}
}
