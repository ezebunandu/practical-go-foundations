package nlp_test

import (
    "fmt"

    "github.com/ardanlabs/nlp"
)

func ExampleTokenize() {
    tokens := nlp.Tokenize("A lil bit of project engineering")
    fmt.Println(tokens)

    // Output:
    // [a lil bit of project engineer]
}