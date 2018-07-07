package sooslib

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var osGetwd = os.Getwd

func cwd() string {
	dir, err := osGetwd()

	check(err)

	return dir
}
