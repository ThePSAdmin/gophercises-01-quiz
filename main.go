package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var (
	totalQuestions = 0
	correctAnswers = 0
)

func main() {
	csvfile := flag.String("csv", "problems.csv", "a csv file")
	limit := flag.Int("limit", 5, "Time limit for each question")
	flag.Parse()

	// Read problems.csv
	file, err := os.Open(*csvfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
problemloop:
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		totalQuestions++

		fmt.Println(row[0], "= ")
		c := make(chan string)
		go func() {
			var input string
			fmt.Scanln(&input)
			c <- input
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-c:
			if answer == row[1] {
				correctAnswers++
			}

		}
	}

	fmt.Println("You answered ", correctAnswers, " correct questions out of ", totalQuestions)
}
