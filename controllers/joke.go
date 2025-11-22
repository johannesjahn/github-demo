package controllers

import (
	"encoding/json"
	"net/http"

	"johannes-jahn.com/demo/joke"
	"johannes-jahn.com/demo/logger"
)

func jokeHandler(jokeService joke.Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logger.Println("jokeHandler called from", r.RemoteAddr)
		joke := jokeService.GetJoke()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(joke); err != nil {
			http.Error(w, "Failed to encode joke", http.StatusInternalServerError)
		}
	}
}
