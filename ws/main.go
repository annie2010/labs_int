// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gopherland/labs_int/ws/internal/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	port                = ":5000"
	defaultReadTimeout  = 10 * time.Second
	defaultWriteTimeout = 10 * time.Second
)

func main() {
	<<!!YOUR_CODE!!>> -- create a gorilla mux and define your route /v1/wc/book/word
	<<!!YOUR_CODE!!>> -- Define your logging middleware

	svc := &http.Server{
		Handler:      r,
		Addr:         port,
		WriteTimeout: defaultWriteTimeout,
		ReadTimeout:  defaultReadTimeout,
	}
	log.Printf("[WordCount] Service listening on port %s", port)
	log.Panic(svc.ListenAndServe())
}
