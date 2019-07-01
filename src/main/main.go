package main

// Параллельное исполнение
// Сделать функцию для параллельного выполнения N заданий.
// Принимает на вход слайс с заданиями `[]func()error`, число заданий
// которые можно выполнять параллельно `N` и максимальное число ошибок после
// которого нужно приостановить обработку. Учесть что задания могу выполняться разное время.

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var now = time.Now()

func sampleFunc() error {
	fmt.Printf("func at %v\n", time.Now().Sub(now))
	return fmt.Errorf("error")
}

func runWorker(tasks chan func() error, errors chan struct{}, maxError int) {
	// endless worker if not errors limit reached or payload exhausted
	for {
		if task, ok := <-tasks; ok {  // if there are some tasks available in the pipeline/channel
			if len(errors) == maxError {  
				// if clean the pipeline if maxErrors were occurred to relieve wg.Wait()
				// i.e. to make sure there are no 'alive' processes awaited by wg.Wait()
				wg.Done()
				continue
			} else if err := task(); err != nil {  // add the channel `errors` if a task is erroneous
				if len(errors) != maxError {
					errors <- struct{}{}
				}
			}
			wg.Done()
		} else {  // if the channel `tasks` has no paylaod anymore
			break
		}
	}
}

func Run (payload []func()error, maxError int, workerCount int) {
	tasks := make(chan func() error, len(payload))
	errors := make(chan struct{}, maxError)

	// start the workers in the background to read and execute from the channel `tasks`
	for i := 0; i < workerCount; i++ {
		go runWorker(tasks, errors, maxError)
	}

	// start pushing the payload to the channel `tasks`
	// such that the workers above would finish the payload
	for _, task := range payload {
		wg.Add(1)
		if len(errors) == maxError {
			break
		}
		tasks <- task
	}
	close(tasks)
	wg.Wait()
}

func main() {
	maxError := 2
	workerCount := 3
	var payload []func()error
	for i := 0; i < 15; i++ {
		payload = append(payload, sampleFunc)
	}
	Run(payload, maxError, workerCount)
}
