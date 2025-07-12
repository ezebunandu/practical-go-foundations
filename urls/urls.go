package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
    urls := []string {
        "https://go.dev",
        "https://grafana.com",
        "https://sans.org/no/such/page",
    }
    start := time.Now()
/*     for _, url := range urls {
        status, err := urlCheck(url)
        if err != nil {
            log.Printf("Error: %q\n", err)
        }
        log.Printf("url: %#v, status: %d\n", url, status)
    } */
    fanOutResult(urls)
    duration := time.Since(start)
    fmt.Printf("total run duration: %v\n", duration)
}

func fanOutResult(urls []string){
    type result struct {
        url string
        status int
        err error
    }
    ch := make(chan result)
    // fanout
    for _, url := range urls {
        go func() {
            r := result{url: url}
            defer func() {ch <- r}()
            r.status, r.err = urlCheck(url)
        }()
    }
    // collect results
    for range urls {
        r := <- ch
        log.Printf("%q: %d (%#v)\n", r.url, r.status, r.err)
    }
    
}

func urlCheck(url string) (int, error) {
    resp, err := http.Get(url)
    if err != nil {
        return 0, err
    }

    return  resp.StatusCode, nil
}