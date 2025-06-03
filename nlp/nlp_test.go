package nlp

import (
	"slices"
	"testing"
)

func TestTokenize(t *testing.T) {
    // setup : call a function
    // teardown: defer/t.Cleanup
    text := "A lil bit of project engineering"
    tokens := Tokenize(text)
    expected := []string{"a", "lil", "bit", "of", "project", "engineering"}
    if !slices.Equal(tokens, expected){
        t.Fatalf("expected %#v, got %#v", expected, tokens)
    }
}