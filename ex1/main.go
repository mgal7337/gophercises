package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

var filepath string

func main() {
	flag.StringVar(&filepath, "path", "./ex1/questions.csv", "path to csv with questions")
	flag.Parse()

	excercises, err := loadQuestionsFromCSVFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var userAnswer string

	for i, v := range excercises {
		question := v[0]
		answer := v[1]

		fmt.Printf("Quetions %d\n%s=", i+1, question)
		fmt.Scanln(&userAnswer)
		if userAnswer == answer {
			fmt.Printf("\nCorrect!\n")
		} else {
			fmt.Printf("\nIncorrect! Correct answer -> %s\n", answer)
		}
	}
}

func loadQuestionsFromCSVFile(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
