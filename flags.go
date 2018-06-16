package main

import "flag"

type flags struct {
	fileName string
}

func getFlags() *flags {
	flg := &flags{}
	flag.StringVar(&flg.fileName, "filename", "problems.csv", "file name to parse questions")
	flag.Parse()
	return flg
}
