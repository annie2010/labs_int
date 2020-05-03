package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "net/http/pprof"

	"github.com/gopherland/ntnx_labs/performance/fib"
)

const httpPort = ":4500"

type (
	result struct {
		Number    int `json:"number"`
		Fibonacci int `json:"fib"`
	}

	results []result
)

func main() {
	http.HandleFunc("/fib", fibHandler)
	log.Printf("Fib is listening on [%s]", httpPort)
	log.Fatal(http.ListenAndServe(httpPort, nil))
}

func fibHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		n   int
	)

	if n, err = strconv.Atoi(r.URL.Query().Get("n")); err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	var res results
	for i := 0; i <= n; i++ {
		res = append(res, result{
			Number:    i,
			Fibonacci: fib.Compute(i),
		})
	}

	buff, err := json.Marshal(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, string(buff))
}
