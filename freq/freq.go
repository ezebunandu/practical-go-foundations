package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)

func main() {
    file, err := os.Open("sherlock.txt")
    if err != nil {
        log.Fatalf("error: %s", err)
    }
    defer file.Close()

    freqs, _ := wordFrequency(file)
    topNwords, err := mostCommonWords(freqs, 10)
    if err != nil {
        log.Fatalf("error: %s", err)
    }

    fmt.Printf("most common n words: %s\n", topNwords)
}

func mostCommonWords(freqs map[string]int, n int) ([]string, error){
    if len(freqs) == 0 {
        return []string{}, fmt.Errorf("empty map")
    }
    words := []wordFreq{}
    for word, freq := range freqs {
        words = append(words, wordFreq{word: word, freq: freq})
    }
    sort.Slice(words, func(i, j int) bool {
        return words[i].freq > words[j].freq
    })
    topN := make([]string, 0, n)
    for _, word := range words[:n]{
        topN = append(topN, word.word)
    }
    return topN, nil
}

type wordFreq struct {
    word string
    freq int
}

func wordFrequency(r io.Reader) (map[string]int, error) {
    s := bufio.NewScanner(r)
    freqs := make(map[string]int) // word -> count
    for s.Scan() {
        words := wordRe.FindAllString(s.Text(), -1)
        for _, w := range words {
            freqs[strings.ToLower(w)]++
        }
    }
    if err := s.Err(); err != nil {
        return nil, err
    }
    return freqs, nil
}