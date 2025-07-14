package stemmer_test

import (
	"fmt"

	"github.com/ardanlabs/nlp/stemmer"
)

func ExampleStem() {
    words := []string{"worked", "working", "works"}
    for _, w := range words {
        fmt.Printf("%s -> %s\n", w, stemmer.Stem(w))
    }

    // Output:
    // worked -> work
    // working -> work
    // works -> work
}