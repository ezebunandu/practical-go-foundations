package main

import (
	"fmt"
	"net/http"
)

// Given a github user login, return name and number of public repos
func main() {
    resp, err := http.Get("https://api.github.com/users/ardanlabs")
    if err != nil {
        fmt.Println("ERROR:", err)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        fmt.Printf("ERROR: bad status - %s\n", resp.Status)
        return
    }
    
}