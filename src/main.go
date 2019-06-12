// Написать функцию находящую максимальный элемент в слайсе
// с произвольными элементами ([]interface{}) с использованием
// пользовательской функции-компаратора.

package main

import (
	"log"
)

func max(items ...int) int {
    maxSoFar := items[0]
    for _, item := range items {
        if item > maxSoFar {
            maxSoFar = item
        }
    }
    return maxSoFar
}

func main() {
	example := []int{-1, -19, -19, -2, -3, -6}
    log.Printf("Max number: %v", max(example...))
    log.Printf("Max number: %v", max(1, -3, 2, 3, 0))
}
