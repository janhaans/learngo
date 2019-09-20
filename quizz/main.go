package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

func ReadCSV(dir string, file string) []Problem {
	var problems []Problem
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	csvFile, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		problems = append(problems, Problem{
			Question: record[0],
			Answer:   record[1],
		})
	}
	return problems
}

func main() {
	problems := ReadCSV("/home/stellar/go/src/github.com/janhaans/learngo/quizz/", "problems.csv")
	reader := bufio.NewReader(os.Stdin)
	var ok, nok int

	for _, problem := range problems {
		fmt.Printf("%s = ", problem.Question)
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		answer = strings.TrimSpace(answer)
		if answer == problem.Answer {
			ok++
		} else {
			nok++
		}
	}

	fmt.Printf("Total number of questions = %d of which you gave %d correct answers\n", ok+nok, ok)
}
