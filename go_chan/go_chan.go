package main

import (
	"fmt"
	"time"
)

func main(){
    go fmt.Println("goroutine")
    fmt.Println("main")

    for i := 0; i < 3; i++{
        go func(n int) {
            fmt.Println(n)
        }(i)
    }
    time.Sleep(10 * time.Millisecond)
}