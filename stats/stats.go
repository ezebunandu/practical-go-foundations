package main

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println(Relu(7))
    fmt.Println(Relu(-1))
    fmt.Println(Relu(1.3))
    fmt.Println(Relu(time.February))

}

/* func ReluInt(i int) int {
    if i < 0 {
        return 0
    }
    return i
}

func ReluFloat64(i float64) float64 {
    if i < 0 {
        return 0
    }
    return i
} */

// T is a "type constraint"
func Relu[T ~int | ~float64] (i T) T {
    if i < 0 {
        return 0
    }
    return i
}