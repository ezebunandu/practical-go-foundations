package main

import (
	"fmt"
	"time"
	// "time"
)

// the worst sort algorithm in the world
func sleepSort(values []int) []int{
    ch := make(chan int, len(values))
    for _, n := range values {
        n := n
        go func(){
            time.Sleep(time.Duration(n) * time.Millisecond)
            ch <- n
        }()
    }
    var res []int
    for range values {
        n := <- ch
        res = append(res, n)
    }
    close(ch)
    return res
}

func main(){
    vals := []int {1, 5, 7, 3, 9}
    res := sleepSort(vals)
    fmt.Printf("Result of sleep sort: %#v\n", res)
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