package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gopherland/labs_int/prometheus/hangman"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const port = ":5000"

var dic *hangman.Dictionary

func init() {
	rand.Seed(time.Now().UnixNano())
	mustLoadDictionary("assets", "words.txt")
}

func main() {
	r := mux.NewRouter()
	m := handlers.LoggingHandler(os.Stdout, r)

	r.Handle("/metrics", promhttp.Handler()).Methods("GET")
	r.Handle("/api/v1/new_game", http.HandlerFunc(newGameHandler)).Methods("GET")
	r.Handle("/api/v1/guess", http.HandlerFunc(guessHandler)).Methods("POST")

	log.Printf("Hangman Listening on port %s...\n", port)
	log.Panic(http.ListenAndServe(port, m))
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Game  *hangman.Game `json:"game"`
		Guess rune          `json:"guess"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	req.Game.Guess(rune(req.Guess))
	raw, err := json.Marshal(req.Game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(raw); err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
	}
}

func newGameHandler(w http.ResponseWriter, r *http.Request) {
	g := hangman.NewGame(pick())

	buff, err := json.Marshal(g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(buff); err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
	}
}

func mustLoadDictionary(dir, file string) {
	var err error
	if dic, err = hangman.NewDictionary(dir, file); err != nil {
		panic(err)
	}
}

func pick() string {
	words := dic.Words()
	return words[rand.Intn(len(words))]
}
