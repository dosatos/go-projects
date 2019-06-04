// Unpacks strings such that "a4bc2d5e" => "aaaabccddddde"
// or "qwe\4\5" => "qwe45", while "45" is an invalid input string

package main

import (
	"fmt"
	"log"
	"strings"
)

func strToInt(digit string) int {
	integer := 0
	order := 1
	for idx := range digit {
		val := int(digit[len(digit)-1-idx] - 48)
		integer += val * order
		order *= 10
	}
	return integer
}

func isDigit(char int32) bool {
	if char >= 48 && char <= 57 {
		return true
	}
	return false
}

func sanityCheck(givenString string) error {
	if isDigit(int32(givenString[0])) {
		return fmt.Errorf("%v is a wrong string", givenString)
	}
	return nil
}

func update(result *strings.Builder, count *strings.Builder, lastNonDigitRune *rune) {
	for i := strToInt(count.String()); i > 0; i-- {
		result.WriteRune(*lastNonDigitRune)
	}
}

func transform(givenString string) string {
	err := sanityCheck(givenString)
	if err != nil {
		return "WRONG INPUT"
	}

	var result strings.Builder
	var count strings.Builder
	var lastNonDigitRune rune
	for _, r := range givenString {
		if !isDigit(r) {
			if count.String() != "" {
				update(&result, &count, &lastNonDigitRune)
				count.Reset()
			} else {
				result.WriteRune(lastNonDigitRune)
			}
			lastNonDigitRune = r
		} else {
			count.WriteRune(r)
		}
	}

	if count.String() != "" {
		update(&result, &count, &lastNonDigitRune)
	} else {
		result.WriteRune(lastNonDigitRune)
	}
	return result.String()
}

func main() {
	givenStrings := []string{
		"he2ы2",
		"he5ы3",
		"he5ы",
		"3he5ы3",
	}
	for _, givenString := range givenStrings {
		res := transform(givenString)
		if res == "WRONG INPUT" {
			log.Println(givenString, res)
		} else {
			log.Printf("%v => %v\n", givenString, transform(givenString))
		}

	}
}
