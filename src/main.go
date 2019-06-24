// Функция логирования Otus
// Задача: написать функцию логирования LogOtusEvent, на вход которой приходят события типа
// HwAccepted (домашняя работа принята) и HwSubmitted (студент отправил дз)
// для этого - спроектировать и реализовать интерфейс OtusEvent.
// Для события HwAccepted мы хотим логирровать дату, айди и грейд,
// для HwSubmitter - дату, id и комментарий, например:

// 2019-01-01 submitted 3456 "please take a look at my homework"
// 2019-01-01 accepted 3456 4
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type HwSubmitted struct {
	Id      int
	Comment string
}

type HwAccepted struct {
	Id    int
	Grade int
}

func (e HwAccepted) String() string {
	return fmt.Sprintf("id: %v, grade: %v", e.Id, e.Grade)
}

func (e HwSubmitted) String() string {
	return fmt.Sprintf("id: %v, comment: '%v'", e.Id, e.Comment)
}

func (e *HwAccepted) logPrefix() string {
	return "accepted"
}

func (e *HwSubmitted) logPrefix() string {
	return "submitted"
}

type OtusEvent interface {
	String() string
	logPrefix() string
}

func LogOtusEvent(event *OtusEvent, file io.Writer) {
	writer := bufio.NewWriter(file)
	line := logMaker(event)
	writer.WriteString(line + "\n")
	writer.Flush()
}

func logMaker(event *OtusEvent) string {
	timePrefix := time.Now().Format("2006/01/02 15:04:05")
	return fmt.Sprintf("%v %v %v", timePrefix, (*event).logPrefix(), (*event).String())
}

func main() {
	events := []OtusEvent{
		&HwSubmitted{Id: 1, Comment: "you are right"},
		&HwAccepted{Id: 1, Grade: 4},
		&HwSubmitted{Id: 2, Comment: "the best"},
		&HwAccepted{Id: 2, Grade: 5},
	}

	// Create or open file
	logFile := "newfile.log"

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		file, err = os.Create(logFile)
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()

	// Log data
	for _, e := range events {
		LogOtusEvent(&e, file)
	}
}
