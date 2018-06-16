package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func presentQuizzes(flg *flags, quizzes [][]string) {

	var input string
	if flg.shuffle {
		shuffle(quizzes)
	}
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

func shuffle(quizzes [][]string) {
	for i := range quizzes {
		j := rand.Intn(i + 1)
		quizzes[i], quizzes[j] = quizzes[j], quizzes[i]
	}
}
