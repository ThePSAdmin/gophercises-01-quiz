package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	totalQuestions = 0
	correctAnswers = 0
)

func main() {
	csvfile := "problems.csv"

	// Read problems.csv
	file, err := os.Open(csvfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		totalQuestions++
		fmt.Println(row[0], "?")
		var input string
		fmt.Scanln(&input)

		if input == row[1] {
			correctAnswers++
		}
	}

	fmt.Println("You answered ", correctAnswers, " correct questions out of ", totalQuestions)
}
