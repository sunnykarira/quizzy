package main

import "flag"

type flags struct {
	fileName string
	shuffle  bool
}

func getFlags() *flags {
	flg := &flags{}
	flag.StringVar(&flg.fileName, "filename", "problems.csv", "file name to parse questions")
	flag.BoolVar(&flg.shuffle, "shuffle", false, "shuffle the questions")
	flag.Parse()
	return flg
}
