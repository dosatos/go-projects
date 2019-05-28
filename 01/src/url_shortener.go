package main

import (
	"fmt"
	"math/rand"
	"path"
	"strconv"
)

const (
	baseURL = "http://short.io/"
)

type MyShortener struct {
	storage map[string]string
}

func (s *MyShortener) Shorten(url string) string {
	var short string
	if s.storage == nil {
		s.storage = make(map[string]string)
	}
	for {
		suffix := strconv.Itoa(rand.Intn(2 ^ 32))
		short = path.Join(baseURL, suffix)
		if _, ok := s.storage[short]; ok {
			fmt.Println("Ouch")
			continue
		}
		s.storage[short] = url
		fmt.Println(s.storage)
		break
	}
	return short
}

func (s *MyShortener) Resolve(url string) string {
	if _, ok := s.storage[url]; ok {
		return s.storage[url]
	}
	fmt.Printf("Cannot resolve for the short url `%v` as it was not shortened before.\n", url)
	return ""
}

func main() {
	urls := []string{"https://www.google.com/", "http://amazon.com", "http://one-more.io"}
	s := new(MyShortener)
	for _, long := range urls {
		fmt.Printf("Shortenning for: %v", long)
		fmt.Printf("Type: %T, Value: %v\n", s, s)
		short := s.Shorten(long)
		fmt.Printf("Short: %v\n", short)
		resolved := s.Resolve(short)
		fmt.Printf("Long: %v\n", resolved)
		fmt.Println()
	}
	s.Resolve("Unknown url")
	fmt.Print(s.storage)
}
