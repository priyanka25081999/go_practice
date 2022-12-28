package main

import (
	"encoding/csv"
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
func problemPuller(fileName string) ([]problem, error) {
	// open the file
	fileObj, err := os.Open(fileName)
	if err == nil {
		// no error, so read the file
		csvR := csv.NewReader(fileObj)
		if line, err := csvR.ReadAll(); err == nil {
			return parseProblem(line), nil
		} else {
			return nil, fmt.Errorf("error in reading csv file")
		}
	} else {
		return nil, fmt.Errorf("error in opening the file %s", fileName)
	}

}

// It parses the slice of string returned by problem puller function
// and divides it into problem slice such that it can contains the actual que and ans format
// It takes lines which are read by problem puller one by one
func parseProblem(lines [][]string) []problem {
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i += 1 {
		r[i] = problem{question: lines[i][0], answer: lines[i][1]}
	}

	return r
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
	correctAns := 0
	timerObj := time.NewTimer(time.Duration(*timer) * time.Second)

	// start looping through the problems
	// make a channel
	ansChannel := make(chan string)
problemLoop:
	for i, p := range prob {
		var ans string
		fmt.Printf("Problem %d: %s = ", i+1, p.question)

		// get the answer and compare it with actual answer
		// let's use goroutine now
		go func() {
			fmt.Scanf("%s", &ans) // get the ans
			ansChannel <- ans
		}()
		select {
		// there are 2 cases here
		// 1st, user completed all the quiz before timer ends
		// 2nd, timer finishes/runs out
		// in both the cases, we want to end our quiz and show the correct answer count to the user
		case <-timerObj.C:
			// this is 2nd case, timer has run out, so we are breaking
			fmt.Println("\nTimer ends!")
			break problemLoop
		case iAns := <-ansChannel:
			if iAns == p.answer {
				correctAns++
			}
			if i == len(prob)-1 {
				close(ansChannel)
			}
		}
	}

	// calculate the result and return it to the user
	fmt.Printf("Your result is %d out of %d\n", correctAns, len(prob))
	fmt.Printf("Press enter to exit")
	<-ansChannel
}
