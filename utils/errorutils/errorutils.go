package errorutils

import (
	"log"
)

func PanicOnError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ExitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
