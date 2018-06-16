package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var score int

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
	presentQuizzes(quizzes)
	_, err = fmt.Fprintf(os.Stdout, "Your total score is %d\n", score)
	if err != nil {
		errorHandler(err)
	}
}

func errorHandler(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
