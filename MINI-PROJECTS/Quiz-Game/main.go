package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "File should be of the form ('Question','Answer')")
	timeLimit := flag.Int("time", 30, "This is the time limit for the quiz")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exitWithError(fmt.Sprintf("There has been a error reading the file: %s", *csvFileName))

	}

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()

	if err != nil {
		exitWithError(fmt.Sprintf("There has been a error parsing lines from the file: %s", *csvFileName))
	}

	problems := parseLines(lines)
	fmt.Println("Welcome to the quiz game!")
	correct := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemLoop:
	for i, prblm := range problems {
		fmt.Printf("Question #%d: %s = \n", i+1, prblm.question)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == prblm.answer {
				fmt.Println("Correct!")
				correct++
			}
		}
	}
	fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
}

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			strings.TrimSpace(line[0]),
			strings.TrimSpace(line[1]),
		}
	}
	return problems
}
func exitWithError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
