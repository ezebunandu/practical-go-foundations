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
    ch := make(chan string) 
    go func() {
        for i := 0; i < 3; i++{
            msg := fmt.Sprintf("message #%d", i+1)
            ch <- msg
        } 
        close(ch)
    }()
    for msg := range ch {
        fmt.Println("got:", msg)
    }
    msg, ok := <- ch // ch is closed
    fmt.Printf("closed: %#v (ok=%v)\n", msg, ok)


    /* channel semantics
    - send & receive will block until opposite operation (*)
    - receive from a closed channel will return the zero value without blocking 
    - send to a closed channel will panic
    - closing a closed channel will panic
    - send/receive to a nil channel will block forever
    */
}