package main

import (
	"log"
	"math/rand"
	"path"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

const (
	baseURL = "http://short.io/"
)

type MyShortener struct {
	storage map[string]string
}

func NewMyShortner() *MyShortener {
	var s MyShortener
	s.storage = make(map[string]string)
	log.Printf("A new shortener has bee initalized: %v\n", s)
	return &s
}

func (s *MyShortener) Shorten(url string) string {
	var short string
	for {
		suffix := strconv.Itoa(rand.Intn(2 ^ 32))
		short = path.Join(baseURL, suffix)
		if _, ok := s.storage[short]; ok {
			continue
		}
		s.storage[short] = url
		break
	}
	return short
}

func (s *MyShortener) Resolve(url string) string {
	if _, ok := s.storage[url]; ok {
		return s.storage[url]
	}
	log.Printf("Cannot resolve for the short url `%v` as it was not shortened before.\n", url)
	return ""
}

func main() {
	urls := []string{"https://www.google.com/", "http://amazon.com", "http://one-more.io"}
	s := NewMyShortner()
	for _, long := range urls {
		log.Printf("Shortenning for: %v\n", long)
		short := s.Shorten(long)
		log.Printf("Short: %v\n", short)
		resolved := s.Resolve(short)
		log.Printf("Long: %v\n\n", resolved)
	}
	s.Resolve("Unknown url")
	log.Printf("Storage: %v", spew.Sdump(s.storage))
}
