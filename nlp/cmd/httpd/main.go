package main

import (
	"encoding/json"
    "expvar"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/ardanlabs/nlp"
	"github.com/ardanlabs/nlp/stemmer"
)

var (
    stemCalls = expvar.NewInt("stem.calls")
)
func main(){
    api := API {
        log: slog.Default().With("app", "nlp"),
    }
    http.HandleFunc("GET /health", api.healthHandler)
    http.HandleFunc("POST /tokenize", api.tokenizeHandler)
    http.HandleFunc("GET /stem/{word}", api.stemHandler)

    addr := ":8080"
    if err := http.ListenAndServe(addr, nil); err != nil {
        fmt.Fprintf(os.Stderr, "error: %s", err)
        os.Exit(1)
    }
}

type tokenizeResponse struct {
    Tokens []string `json:"tokens"`
}

type API struct {
    log *slog.Logger
}

func (a *API) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
    data, err := io.ReadAll(r.Body)
    if err != nil {
        a.log.Error("read", "error", err, "remote", r.RemoteAddr)
        http.Error(w, "can't read data", http.StatusBadRequest)
        return
    }
    text := string(data)
    if len(text) == 0 {
        a.log.Error("read", "error", "empty request")
        http.Error(w, "empty request", http.StatusBadRequest)
        return
    }
    w.Header().Set("content-type", "application/json")
    tokens := nlp.Tokenize(text)
    err = json.NewEncoder(w).Encode(tokenizeResponse{Tokens: tokens})
    if err != nil {
        a.log.Error("read", "error", "internal server error")
        http.Error(w, fmt.Sprintf("error building response, %v", err), http.StatusInternalServerError)
        return 
    }
}

func (a *API) stemHandler(w http.ResponseWriter, r *http.Request) {
    stemCalls.Add(1)
    word := r.PathValue("word")
    a.log.Info("stem", "word", word)
    fmt.Fprintln(w, stemmer.Stem(word))
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
    if err := health(); err != nil {
        a.log.Error("health", "error", err)
        http.Error(w, "health check failed", http.StatusInternalServerError)
        return
    }
    fmt.Fprintln(w, "OK")
}

func health() error {
    // TODO: Actual health check
    return  nil
}