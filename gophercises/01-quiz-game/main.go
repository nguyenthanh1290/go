package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type options struct {
	csvFilename string
	timeLimit   time.Duration
}

type problem struct {
	q string
	a string
}

func main() {
	userOptions := readInput()
	problems := loadProblems(userOptions.csvFilename)
	timer := time.NewTimer(userOptions.timeLimit * time.Second)
	correct := 0

problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)

		answerChannel := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerChannel <- ans
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case ans := <-answerChannel:
			if ans == p.a {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
	fmt.Println("Bye...")
}

func readInput() options {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit fo the quiz in seconds")
	flag.Parse()

	return options{*csvFilename, time.Duration(*timeLimit)}
}

func loadProblems(fileName string) []problem {
	file, err := os.Open(fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", fileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	return parseProblems(lines)
}

func parseProblems(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
