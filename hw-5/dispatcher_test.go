package main

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var TestCases = []struct {
	testCaseID       string         // Name of testcase
	numberOfTasks    int            // Number of tasks to run
	numberOfWorkers  int            // Number of workers to handle tasks in parallel
	numberOfErrors   int            // Errors threshold
	errorProbability int            // the probability of error appearance 10 - 0 errors. >11 no errors
	expectedResult   error          // Expected Error code from the Dispatcher
	tasks            []func() error // Slice of task generated randomly
}{
	{
		testCaseID:       "Basic test 1, no error expected. Number of tasks are equals to the number of workers",
		numberOfTasks:    4,
		numberOfWorkers:  4,
		numberOfErrors:   2,
		errorProbability: 11,
		expectedResult:   nil,
	},
	{
		testCaseID:       "Basic test 2, no error expected. Number of task in 5 times higher then number of workers",
		numberOfTasks:    20,
		numberOfWorkers:  4,
		numberOfErrors:   2,
		errorProbability: 11,
		expectedResult:   nil,
	},
	{
		testCaseID:       "Basic test 3, no error expected. Number of task in 20 times higher then number of workers",
		numberOfTasks:    20,
		numberOfWorkers:  1,
		numberOfErrors:   2,
		errorProbability: 11,
		expectedResult:   nil,
	},
	{
		testCaseID:       "Basic test 4, no error expected. Number of tasks is less in 2  times then number of workers",
		numberOfTasks:    8,
		numberOfWorkers:  16,
		numberOfErrors:   2,
		errorProbability: 11,
		expectedResult:   nil,
	},
	{
		testCaseID:       "Basic test 5, expected 2 errors. Number of task in 5 times higher then number of workers",
		numberOfTasks:    20,
		numberOfWorkers:  4,
		numberOfErrors:   2,
		errorProbability: 5,
		expectedResult:   ErrTaskError,
	},
	{
		testCaseID:       "Basic test 6, Two errors expected. Number of task in 20 times higher then number of workers",
		numberOfTasks:    20,
		numberOfWorkers:  1,
		numberOfErrors:   2,
		errorProbability: 5,
		expectedResult:   ErrTaskError,
	},
	{
		testCaseID:       "Basic test 7, Two errors expected. Number of task in 5 times higher then number of workers",
		numberOfTasks:    8,
		numberOfWorkers:  4,
		numberOfErrors:   2,
		errorProbability: 5,
		expectedResult:   ErrTaskError,
	},
	//	{
	//		testCaseID:       "Basic test 8, expected 5 error. Number of tasks are less then number of workers",
	//		numberOfTasks:    4,
	//		numberOfWorkers:  8,
	//		numberOfErrors:   1,
	//		errorProbability: 1,
	//		expectedResult:   ErrTaskError,
	//	},
}

func TestDispatcher(t *testing.T) {
	tasksBuilder()
	for _, test := range TestCases {
		err := Dispatcher(test.tasks, test.numberOfWorkers, test.numberOfErrors)
		if err != test.expectedResult {
			t.Errorf("Expected %s but received %v", test.expectedResult, err)
		}
		continue
	}
}
func tasksBuilder() {
	for i, test := range TestCases {
		tasks := make([]func() error, 0) // Prepare a slice of tasks
		for j := 0; j < test.numberOfTasks; j++ {
			j := j
			ep := test.errorProbability
			task := func() error {
				fmt.Println("Startingin task ID=", j)
				start := time.Now()
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
				if rand.Intn(10) > ep {
					fmt.Println("Time since start to get an error id", j, time.Since(start))
					return errors.New("Task aborted with error")
				}
				fmt.Println("Task finished successfully id=", j, time.Since(start))
				return nil
			}
			tasks = append(tasks, task)
		}
		//test.tasks = tasks
		TestCases[i].tasks = tasks
		//fmt.Println(test.tasks)
	}
}
