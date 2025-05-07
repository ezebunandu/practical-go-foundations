package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
    banner("Go", 6)
    banner("G♡", 6)
}

func banner(text string, width int) {
    padding := (width - utf8.RuneCountInString(text))/2
    fmt.Print(strings.Repeat(" ", padding))
    fmt.Println(text)
    fmt.Println(strings.Repeat("-", width))
}