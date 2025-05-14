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

    m, err := NewMatrix[float64](10, 3)
    if err != nil {
        fmt.Println("ERROR:", err)
    }
    fmt.Println("m:", m)
    fmt.Println(m.At(3, 2))
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

type Matrix [T Number] struct {
    Rows int
    Cols int
    data []T
}

func NewMatrix[T Number] (rows, cols int) (*Matrix[T], error) {
    if rows <= 0 || cols <= 0 {
        return nil, fmt.Errorf("bad dimensions: %d/%d", rows, cols)
    }
    m := Matrix[T] {
        Rows: rows,
        Cols: cols,
        data: make([]T, rows * cols),
    }
    return &m, nil
}

func (m *Matrix[T]) At (row, col int) T {
    i := (row * m.Cols) + col
    return m.data[i]
}
type Number interface {
    ~int | ~float64
}

// T is a "type constraint"
func Relu[T Number] (i T) T {
    if i < 0 {
        return 0
    }
    return i
}