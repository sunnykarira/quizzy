package main

import (
	"fmt"
	"os"
	"strings"
)

func presentQuizzes(quizzes [][]string) {
	var input string
	for _, quiz := range quizzes {
		question := strings.TrimSpace(quiz[0])
		answer := strings.TrimSpace(quiz[1])
		_, err := fmt.Fprintf(os.Stdout, "The problem is %s\n", question)
		if err != nil {
			errorHandler(err)
		}
		fmt.Scanln(&input)
		if strings.TrimSpace(input) == answer {
			score++
			fmt.Fprint(os.Stdout, "Correct Answer\n")
		}
	}
}
