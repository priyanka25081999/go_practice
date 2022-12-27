package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

// This pulls the problem from quiz.csv file
// and return the slice of string which contains the que and ans data
// otherwise it returns an error
func problemPuller(fileName string) ([]string, error) {

}

// It parses the slice of string returned by problem puller function
// and divides it into problem slice such that it can contains the actual que and ans format
// It takes lines which are read by problem puller one by one
func parseProblem(lines [][]string) []problem {

}

func main() {
	file := flag.String("f", "quiz.csv", "path of the file")

	// set the timer for the quiz
	timer := flag.Int("t", 30, "timer for the quiz")

	// parse the file and timer flag
	flag.Parse()

	// pull the problem from file
	prob, err := problemPuller(*file)
	if err != nil {
		fmt.Println("Error occured, error:", err)
		os.Exit(0)
	}

	// initialize the correct ans and timer variable
	correctAnd := 0
	timerObj := time.NewTimer(time.Duration(*timer) * time.Second)

}
