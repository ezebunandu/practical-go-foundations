package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ardanlabs/nlp"
	"github.com/ardanlabs/nlp/stemmer"
)

func main(){
    http.HandleFunc("GET /health", healthHandler)
    http.HandleFunc("POST /tokenize", tokenizeHandler)
    http.HandleFunc("GET /stem/{word}", stemHandler)

    addr := ":8080"
    if err := http.ListenAndServe(addr, nil); err != nil {
        fmt.Fprintf(os.Stderr, "error: %s", err)
        os.Exit(1)
    }
}

type tokenizeResponse struct {
    Tokens []string `json:"tokens"`
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
    data, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "can't read data", http.StatusBadRequest)
        return
    }
    text := string(data)
    if len(text) == 0 {
        http.Error(w, "empty request", http.StatusBadRequest)
        return
    }
    w.Header().Set("content-type", "application/json")
    tokens := nlp.Tokenize(text)
    err = json.NewEncoder(w).Encode(tokenizeResponse{Tokens: tokens})
    if err != nil {
        http.Error(w, fmt.Sprintf("error building response, %v", err), http.StatusInternalServerError)
        return 
    }
}

func stemHandler(w http.ResponseWriter, r *http.Request) {
    word := r.PathValue("word")
    fmt.Fprintln(w, stemmer.Stem(word))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    if err := health(); err != nil {
        http.Error(w, "health check failed", http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, "OK")
}

func health() error {
    // TODO: Actual health check
    return  nil
}