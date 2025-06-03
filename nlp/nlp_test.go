package nlp

import (
	"os"
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

func TestTokenizeTable(t *testing.T){
    var cases = []struct{
        text string
        tokens []string
    }{
        {"Who's on first?", []string{"who", "s", "on", "first"},},
        {"A lil bit of project engineering", []string{"a", "lil", "bit", "of", "project", "engineering"},},
        {"", nil},
    }
    for _, tc := range cases {
        t.Run(tc.text, func(t *testing.T) {
            tokens := Tokenize(tc.text)
            if !slices.Equal(tc.tokens, tokens) {
                t.Fatalf("expected %#v, got %#v", tc.tokens, tokens)
            }
        })
    }
}

/*
Selecting tests:
- "run" flag: regexp
- build tags (//go:build comment)
- environment variables
*/

var inCI = os.Getenv("CI") != ""

func TestInCI(t *testing.T){
    if !inCI {
        t.Skip("not in CI")
    }
}