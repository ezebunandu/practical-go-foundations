package main

import (
	"fmt"
	"sync"
	"time"
)

type Payment struct {
	From   string
	To     string
	Amount float64 // USD
	once   sync.Once
}

func (p *Payment) Process() {
	t := time.Now()
	// wrap the call to p.process in an anonymous function so it can be passed
	// to once.Do
	// which will only accept a function that takes no parameters
	// hence reducing the arity
	p.once.Do(func() { p.process(t) })
}

// each payment should only be processed once
func (p *Payment) process(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("[%s] %s -> %.2f -> %s\n", ts, p.From, p.Amount, p.To)
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
