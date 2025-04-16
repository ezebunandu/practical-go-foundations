package main

import (
	"fmt"
	"sync"
)

type Payment struct {
	From   string
	To     string
	Amount float64 // USD
	once   sync.Once
}

func (p *Payment) Process() {
    p.once.Do(p.process)
}

// each payment should only be processed once
func (p *Payment) process() {
	fmt.Printf("%s -> %.2f -> %s\n", p.From, p.Amount, p.To)
}

func main() {
	p := Payment{
		From:   "Wile. E. Coyote",
		To:     "ACME Corp",
		Amount: 123.34,
	}
	p.Process()
	p.Process()
}
