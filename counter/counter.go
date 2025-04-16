package main

import (
	"fmt"
	"sync"
)

func main(){
    var mu sync.Mutex // put the mutex on top of the variable to guard
    count := 0

    const n = 10
    var wg sync.WaitGroup
    wg.Add(n)
    for i := 0; i < n; i++ {
        go func ()  {
            defer wg.Done()
            for j := 0; j < 10_000; j ++ {
                mu.Lock()
                count++
                mu.Unlock()
            }
        }()
    }
    wg.Wait()
    fmt.Println(count)
}