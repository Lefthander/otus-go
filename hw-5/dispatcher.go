// Sergey Olisov (c) 2019
// Lesson 5
// Workers pool
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

// TODO:
// 1. Simple workload function
// 2. Run pool of such functions
// 3. Manage errors from the functions

// structure to hold the Errors Counter
type errorsCnt struct {
	sync.RWMutex
	errorCounter   int // Error counter
	errorThreshold int // Errors threshold
}

// Inc does the increment of errorCounter in threadsafe way
func (ec *errorsCnt) Inc() {
	ec.Lock()
	ec.errorCounter++
	ec.Unlock()
}

// isThreshouldReached check the actual value of errorCounter against defined threshould if greater or equal returns true
func (ec *errorsCnt) isThresholdReached() bool {
	ec.RLock()
	defer ec.RUnlock()
	return ec.errorCounter == ec.errorThreshold
}

// Get returns a current value of errorCounter
func (ec *errorsCnt) Get() int {
	ec.RLock()
	defer ec.RUnlock()
	return ec.errorCounter
}

// ErrTaskError indicate that worker abnormaly terminated due to faulure of processing task
var ErrTaskError = errors.New("Worker aborted due to task error")

func task() error {
	fmt.Println("Starting task...")
	start := time.Now()
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	if rand.Intn(10) >= 5 {
		fmt.Println("Time since start to get an error", time.Since(start))
		return errors.New("Task has an error durinng execution")
	}
	fmt.Println("Ending succesfully... Time since start", time.Since(start))
	return nil
}

// Dispatcher runs a pool of workers (numberOfTasks) and assign the tasks to them from slice tasks
func Dispatcher(tasks []func() error, numberOfTasks int, numberOfErrors int) error {

	workersChannel := make(chan func() error, numberOfTasks) // Channel to assign a tasks to workers
	errorsChannel := make(chan error, numberOfTasks)         // Channel to transfer errors from the tasks
	defer close(errorsChannel)

	shutdownChannel := make(chan bool) // Channel for a shutdown of all workers

	errCnt := &errorsCnt{errorThreshold: numberOfErrors} // Structure to manage number of Errors, reported by workers

	// Managing errors channel...
	go func() {
		for err := range errorsChannel {
			errCnt.Inc() // Increment Error Counter
			fmt.Println("Detected an error", err)
			if errCnt.isThresholdReached() {
				fmt.Println("Threshold number of errors has reached. Aborting...")
				close(shutdownChannel)
				return
			}
		}
	}()

	eg := errgroup.Group{}
	for i := 0; i < numberOfTasks; i++ {
		i := i // hide loop variable
		eg.Go(func() error {
			for task := range workersChannel {
				select {
				case <-shutdownChannel:
					fmt.Println("Worker aborted...", i)
					return ErrTaskError
				default:
					fmt.Println("Worker started", i)
					if err := task(); err != nil {
						errorsChannel <- err
					}
					fmt.Println("Worked stopped", i)
				}
			}
			fmt.Println("Worker finished...", i)
			return nil
		})
	}
	for _, task := range tasks {
		select {
		case <-shutdownChannel:
			break
		default:
			workersChannel <- task
		}
	}
	close(workersChannel)

	err := eg.Wait()
	if err != nil {
		return err
	} else if errCnt.Get() != 0 {
		return ErrTaskError
	}
	return nil
}

func main() {

	tt := []func() error{task, task, task, task}
	fmt.Println("Dispatcher Started...")
	err := Dispatcher(tt, 2, 1)
	if err != nil {
		fmt.Println("Dispatcher finished with errors", err)
	}
	fmt.Println("Dispatcher ended...", err)
}
