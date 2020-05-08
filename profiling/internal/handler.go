package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gopherland/labs_int/profiling/internal/fib"
)

type Result struct {
	Number    int `json:"n"`
	Fibonacci int `json:"fib"`
}

type Results []Result

func FibHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		n   int
	)
	if n, err = strconv.Atoi(r.URL.Query().Get("n")); err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
		return
	}

	var res Results
	for i := 0; i <= n; i++ {
		res = append(res, Result{Number: i, Fibonacci: fib.Compute(i)})
	}

	buff, err := json.Marshal(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(buff); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
