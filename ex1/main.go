package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var filepath string

func main() {
	flag.StringVar(&filepath, "path", "./ex1/questions.csv", "path to csv with questions")
	timeoutVal := flag.Int("time", 5, "time to answer")
	flag.Parse()

	excercises, err := loadQuestionsFromCSVFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	timeout := time.Duration(*timeoutVal) * time.Second

	for i, v := range excercises {
		questionPrinter(v[0], i+1)
		userAnswerChan := make(chan string)
		go getUserAnswer(userAnswerChan)
		select {
		case ans := <-userAnswerChan:
			handleAnswer(v[1], ans)
		case <-time.After(timeout):
			fmt.Println("Timeout!")
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

func getUserAnswer(ret chan<- string) {
	var answer string
	fmt.Scanln(&answer)
	ret <- strings.TrimSpace(strings.ToLower(answer))
}

func questionPrinter(question string, number int) {
	fmt.Printf("#%d: Next question: %s= ? ", number, question)
}

func handleAnswer(correct, provided string) {
	if correct != provided {
		fmt.Printf("Wrong answer!\nProvided: %s, Expected: %s\n", provided, correct)
	} else {
		fmt.Println("Correct!")
	}
}
