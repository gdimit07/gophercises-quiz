package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var limitFlag int
var csvFlag string

func init() {

	flag.IntVar(&limitFlag, "limit", 30, "the time limit for the quiz in seconds")
	flag.StringVar(&csvFlag, "csv", "problems.csv", "a csv file in the format of 'question,answer'")
}

func main() {

	var correctAnswers int
	var wrongAnswers int
	var usersAnswer string

	flag.Parse()

	//read data from the csv file - data variable is of type byte slice
	data, err := os.ReadFile(csvFlag)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	questionsCSV := string(data)

	r := csv.NewReader(strings.NewReader(questionsCSV))

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v:", record[0])
		fmt.Scan(&usersAnswer)

		if usersAnswer == record[1] {
			correctAnswers++
		} else {
			wrongAnswers++
		}

	}

	fmt.Println(correctAnswers, wrongAnswers)

}
