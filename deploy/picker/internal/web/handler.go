// © 2020 Imhotep Software LLC. All rights reserved.

// Package web rpresents a web service.
package web

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gopherland/ntnx_labs/deploy/picker/internal/dictionary"
)

// Handler represents a dictionary handler
type Handler struct {
	assetDir string
	words    dictionary.WordList
}

// Response greeting message format.
type Response struct {
	AssetDir, Dictionary string
	WordCount            int
}

// New returns a new ditionary handler.
func New(dir string) *Handler {
	return &Handler{assetDir: dir}
}

// LoadHandler loads a dictionary in memory.
func (h *Handler) LoadHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		http.Error(w, "you must provide a dictionary name", http.StatusExpectationFailed)
		return
	}

	var err error
	h.words, err = dictionary.Load(h.assetDir, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	resp := Response{
		AssetDir:   "assets",
		Dictionary: name,
		WordCount:  len(h.words),
	}
	buff, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(buff))
}

// PickHandler pick out a random word.
func (h *Handler) PickHandler(w http.ResponseWriter, r *http.Request) {
	if len(h.words) == 0 {
		http.Error(w, "no dictionary loaded", http.StatusInternalServerError)
		return
	}
	word := h.words[rand.Intn(len(h.words))]

	resp := struct {
		Word string
	}{
		Word: word,
	}
	buff, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(buff))
}
