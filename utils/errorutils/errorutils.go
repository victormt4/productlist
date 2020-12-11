package errorutils

import (
	"fmt"
	"os"
)

func PanicOnError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ExitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
