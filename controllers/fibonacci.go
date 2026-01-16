package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"johannes-jahn.com/demo/fibonacci"
	"johannes-jahn.com/demo/logger"
)

func fibonacciHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Println("fibonacciHandler called from", r.RemoteAddr)

		nStr := r.URL.Query().Get("n")
		n := 10 // default
		if nStr != "" {
			var err error
			n, err = strconv.Atoi(nStr)
			if err != nil || n < 0 {
				http.Error(w, "Invalid parameter 'n'", http.StatusBadRequest)
				return
			}
		}

		if n > 90 {
			// Limit safely to avoid overflow and abuse
			http.Error(w, "Parameter 'n' too large (max 90)", http.StatusBadRequest)
			return
		}

		series := fibonacci.Calculate(n)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(series); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	}
}
