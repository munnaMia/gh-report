package utils

import "log"

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func PrintFatal(str string) {
	log.Fatal(str)
}