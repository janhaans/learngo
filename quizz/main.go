package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var dirName = flag.String("d", os.Getenv("HOME"), "Directory where csv file is placed")
var fileName = flag.String("f", "problems.csv", "Name of csv file")
var timer = flag.Int("t", 30, "Set timer duration")

var wg sync.WaitGroup

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
	flag.Parse()
	problems := ReadCSV(*dirName, *fileName)
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Press a Enter to start")
	_, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	var ok, nok int

	wg.Add(1)

	go func() {
		defer wg.Done()
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
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(*timer) * time.Second)
	}()

	wg.Wait()

	fmt.Printf("\nTotal number of questions = %d of which you gave %d correct answers\n", len(problems), ok)
}
