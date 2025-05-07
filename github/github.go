package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Given a github user login, return name and number of public repos
func main() {
    fmt.Println(UserInfo("ezebunandu"))
}
func demo () {
    resp, err := http.Get("https://api.github.com/users/ardanlabs")
    if err != nil {
        fmt.Println("ERROR:", err)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        fmt.Printf("ERROR: bad status - %s\n", resp.Status)
        return
    }
    ctype := resp.Header.Get("Content-Type") 
    fmt.Println("content-type:", ctype)
    var reply struct {
        Name string `json:"name"`
        PublicRepos int `json:"public_repos"`
    }
    dec := json.NewDecoder(resp.Body)
    if err := dec.Decode(&reply); err != nil {
        fmt.Println("ERROR:", err)
        return
    }
    // io.Copy(os.Stdout, resp.Body)
    fmt.Println(reply.Name, reply.PublicRepos)
}

func UserInfo(user string) (string, int, error){
    url := "https://api.github.com/users/" + user
    resp, err := http.Get(url)
    if err != nil {
        return "", 0, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return "", 0, fmt.Errorf("%q - bad status: %s", url, resp.Status)
    }
    return parseResponse(resp.Body)
}

func parseResponse(r io.Reader) (string, int, error) {
    var reply struct {
        Name string `json:"name"`
        PublicRepos int `json:"public_repos"`
    }
    dec := json.NewDecoder(r)
    if err := dec.Decode(&reply); err != nil {
        return "", 0, err
    }
    // io.Copy(os.Stdout, resp.Body)
    return reply.Name, reply.PublicRepos, nil
}