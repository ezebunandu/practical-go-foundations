package nlp

import (
	"os"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"

	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	// setup : call a function
	// teardown: defer/t.Cleanup
	text := "A lil bit of project engineering"
	tokens := Tokenize(text)
	expected := []string{"a", "lil", "bit", "of", "project", "engineering"}

	require.Equal(t, expected, tokens)
	// Before testify
	/*     if !slices.Equal(tokens, expected){
	       t.Fatalf("expected %#v, got %#v", expected, tokens)
	   } */
}

func TestTokenizeTable(t *testing.T) {
	cases := loadTokenizeCase(t)
	for _, tc := range cases {
		name := tc.Name
		if name == "" {
			name = tc.Text
		}
		t.Run(name, func(t *testing.T) {
			tokens := Tokenize(tc.Text)
            if tokens == nil {
                tokens = []string{}
            }
			t.Logf("Text: %q, Got: %#v, Expected: %#v", tc.Text, tokens, tc.Tokens)
			require.Equal(t, tc.Tokens, tokens)
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

func TestInCI(t *testing.T) {
	if !inCI {
		t.Skip("not in CI")
	}
}

type tokCase struct {
	Text   string
	Tokens []string
	Name   string
}

func loadTokenizeCase(t *testing.T) []tokCase {
	file, err := os.Open("testdata/tokenize_cases.toml")
	require.NoError(t, err)
	defer file.Close()

	var data struct {
		Cases []tokCase `toml:"cases"`
	}
	dec := toml.NewDecoder(file)
	_, err = dec.Decode(&data)
	require.NoError(t, err)
	t.Logf("Loaded cases: %+v", data.Cases)
	return data.Cases
}

func FuzzTokenizer(f *testing.F){
    f.Add("")
    fn := func(t *testing.T, text string) {
        tokens := Tokenize(text)
        lText := strings.ToLower(text)
        for _, tok := range tokens {
            require.Contains(t, lText, tok)
        }
    }
    f.Fuzz(fn)
}