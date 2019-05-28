package main

import (
    "fmt"
    "math/rand"
    "strconv"
)


const (
    base_url = "http://short.io/" 
)


func save_to_db(url string, hashed_url string, database map[string]string) {
    database[hashed_url] = url
}


func hash_url(url string) string {
    hashed_url := strconv.FormatInt(int64(rand.Int()), 10)
    fmt.Printf("INFO: %v hashed as %v\n", url, hashed_url)
    return hashed_url
}


func get_short_suffix(url string, database map[string]string) string {
    // there is probably a way to hash it properly
    // random value is used for simplicity
    hashed_url := hash_url(url)
    // check if exists in the db already
    if _, ok := database[hashed_url]; !ok {
        fmt.Printf("INFO: `%v` for %v saved to the db\n", hashed_url, url)
        save_to_db(url, hashed_url, database)
    }
    return hashed_url
}

func parse_key(url string) string {
    start_pos := len(base_url)
    key := url[start_pos:]
    return key
}


func get_from_db(short_url string, database map[string]string) (string, bool) {
    hashed_url := parse_key(short_url)
    original_url, ok := database[hashed_url]
    if ok {
        return original_url, true
    }
    return "", false
}


func get_long(url string, database map[string]string) (string, bool) {
    return get_from_db(url, database)
}


func main() {
    fmt.Println("INFO: Started!")
    urls := []string{"https://www.google.com", "http://amazon.com"}
    database := make(map[string]string)
    for _, url := range urls {
        fmt.Printf("===== Started for: %v ===== \n", url)
        short_url := base_url + get_short_suffix(url, database)
        fmt.Printf("INFO: Original URL: %v\n", url)
        fmt.Printf("INFO: Shortened URL: %v\n", short_url)
        long_url, ok := get_long(short_url, database)
        if ok {
            fmt.Printf("INFO: Original from db URL: `%v`\n", long_url)
        } else {
            fmt.Printf("WARNING: the pair for %v not found!\n", short_url)
        }
    }
    fmt.Printf("INFO: DB state: %v\n", database)
}
