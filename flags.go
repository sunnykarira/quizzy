package main

import "flag"

type flags struct {
	fileName string
	shuffle  bool
	timer    int
}

func getFlags() *flags {
	flg := &flags{}
	flag.StringVar(&flg.fileName, "filename", "problems.csv", "file name to parse questions")
	flag.BoolVar(&flg.shuffle, "shuffle", false, "shuffle the questions")
	flag.IntVar(&flg.timer, "timer", 30, "default expire timer")
	flag.Parse()
	if !checkFlags(flg) {
		return &flags{
			fileName: "problems.csv",
			timer:    30,
		}
	}
	return flg
}

func checkFlags(flg *flags) bool {
	if flg.timer <= 0 {
		return false
	}
	return true
}
