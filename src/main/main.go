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

func Run (payload []func()error, maxError int, workerCount int) {
	tasks := make(chan func() error, len(payload))
	errors := make(chan struct{}, maxError)

	// kick-off a process in the background to read and execute from the tasks
	for i := 0; i < workerCount; i++ {
		go func() {
			// endless worker if not errors limit reached or payload exhausted
			for {
				if task, ok := <-tasks; ok {
					if len(errors) == maxError {
						wg.Done()
						continue
					} else if err := task(); err != nil {
						if len(errors) != maxError {
							errors <- struct{}{}
						}
					}
					wg.Done()
				} else {
					break
				}
			}
		}()
	}

	// start pushing the payload to the tasks
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
	maxError := 1
	workerCount := 2
	var payload []func()error
	for i := 0; i < 15; i++ {
		payload = append(payload, sampleFunc)
	}
	Run(payload, maxError, workerCount)
}
