// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

// Package web rpresents a web service.
package web

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/gopherland/labs_int/picker_svc/internal/dictionary"
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

// New returns a new dictionary handler.
func New(dir string) *Handler {
	return &Handler{assetDir: dir}
}

// LoadHandler loads a dictionary in memory.
func (h *Handler) LoadHandler(w http.ResponseWriter, r *http.Request) {
	dic := r.URL.Query().Get("dictionary")
	if len(dic) == 0 {
		http.Error(w, "you must provide a dictionary name", http.StatusExpectationFailed)
		return
	}

	var err error
	h.words, err = dictionary.Load(h.assetDir, dic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	resp := Response{
		AssetDir:   "assets",
		Dictionary: dic,
		WordCount:  len(h.words),
	}
	buff, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err = w.Write(buff); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
	if _, err := w.Write(buff); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
