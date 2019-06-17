// Написать функцию находящую максимальный элемент в слайсе
// с произвольными элементами ([]interface{}) с использованием
// пользовательской функции-компаратора.

package main

import (
	"log"
)

func max(slice []interface{}, comparator func(slice []interface{}, a, b int) bool) interface{} {
	if len(slice) == 0 {
		return nil
	}

	maxValueIndex := 0
	for idx := range slice {
		if comparator(slice, maxValueIndex, idx) {
			maxValueIndex = idx
		}
	}
	return slice[maxValueIndex]
}

type Person struct {
	age  int
	name string
}

type Salary struct {
	employee Person
	amount   int
}

func runForAge(people []interface{}) {
	// the oldest person
	comparator := func(slice []interface{}, a, b int) bool {
		return slice[b].(Person).age > slice[a].(Person).age
	}

	want := people[1]
	log.Printf("Is equal: %v <== want: `%v`, have: `%v`", want == max(people, comparator), want, max(people, comparator))
}

func runForSalaries(people []interface{}) {
	// the most paid person
	salaries := []interface{}{
		Salary{people[0].(Person), 120},
		Salary{people[1].(Person), 210},
		Salary{people[2].(Person), 1113},
	}

	comparator := func(slice []interface{}, a, b int) bool {
		return slice[b].(Salary).amount > slice[a].(Salary).amount
	}

	want := salaries[2]
	log.Printf("Is equal: %v <== want: `%v`, have: `%v`", want == max(salaries, comparator), want, max(salaries, comparator))
}

func runForOldestEmployeeFromSalaries(people []interface{}) {
	// the oldest employee
	salaries := []interface{}{
		Salary{people[0].(Person), 120},
		Salary{people[1].(Person), 210},
		Salary{people[2].(Person), 1113},
	}

	comparator := func(slice []interface{}, a, b int) bool {
		return slice[b].(Salary).employee.age > slice[a].(Salary).employee.age
	}

	want := salaries[1]
	log.Printf("Is equal: %v <== want: `%v`, have: `%v`", want == max(salaries, comparator), want, max(salaries, comparator))
}

func main() {
	people := []interface{}{
		Person{10, "Alex"},
		Person{20, "John"},
		Person{13, "Maria"},
	}
	runForAge(people)
	runForSalaries(people)
	runForOldestEmployeeFromSalaries(people)
}
