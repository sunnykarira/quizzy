package main

import (
	"encoding/csv"
	"strings"
)

func read(in string) [][]string {

	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ','
	quizzes, err := r.ReadAll()
	if err != nil {
		errorHandler(err)
	}
	return quizzes
}
