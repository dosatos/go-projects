package main

import (
	"log"
	"regexp"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func count(text string) (result []string) {
	text = strings.ToLower(text)
	r, _ := regexp.Compile(`(\w+)`)
	storage := make(map[string]int)
	for _, word := range r.FindAllString(text, -1) {
		storage[word] += 1
	}

	// Transforming into a wordCount struct
	var countedWords []wordCount
	for word, count := range storage {
		countedWords = append(countedWords, wordCount{
			word: word, count: count,
		})
	}

	// O(nlog(n)) ???
	sort.Slice(countedWords, func(a, b int) bool {
		return countedWords[a].count > countedWords[b].count
	})

	// fmt.Println(countedWords)

	// O(1) - space / time
	for i, wc := range countedWords {
		if i == 10 {
			break
		}
		result = append(result, wc.word)
	}
	return result
}

func main() {
	log.Println(count("Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."))
	log.Println(count("Lorem Ipsum is simply dummy text of the printing and the"))
}
