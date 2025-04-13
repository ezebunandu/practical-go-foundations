package main

import (
	"fmt"
	"sort"
)

func main() {
	var s3 []int
	for i := 0; i < 1_000; i++ {
		s3 = appendInt(s3, i)
	}
	fmt.Println("s3", len(s3), cap(s3))
	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"})) // [A B C D E]
    vs := []float64{2, 1, 3, 4}
    fmt.Println(median(vs))
}

func median(values []float64) float64 {
    sort.Float64s(values)
    i := len(values) / 2 
    if len(values)%2 == 1 {
        return values[i]
    }
    v := (values[i-1] + values[i]) / 2 
    return v
}

func concat(s1, s2 []string) []string {
    s := make([]string, len(s1) + len(s2))
    copy(s, s1)
    copy(s[len(s1):], s2)
    return s
}

func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) < cap(s) {
		s = s[:len(s)+1]
	} else {
		fmt.Printf("reallocate: %d->%d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}
	s[i] = v
	return s
}
