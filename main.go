package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Get Flags
	flg := getFlags()
	//Read the file
	f, err := ioutil.ReadFile(flg.fileName)
	if err != nil {
		errorHandler(err)
	}
	quizzes := read(string(f))
	if quizzes == nil {
		errorHandler(errors.New("No Quizzes"))
	}

}

func errorHandler(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
