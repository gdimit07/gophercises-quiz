package main

import (
	"flag"
)

var limitFlag int
var csvFlag string

func init() {

	flag.IntVar(&limitFlag, "limit", 30, "the time limit for the quiz in seconds")
	flag.StringVar(&csvFlag, "csv", "problems.csv", "a csv file in the format of 'question,answer'")
}

func main() {

	flag.Parse()

}
