package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"johannes-jahn.com/demo/joke"
)

func jokeHandler(jokeService joke.Service) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Default().Println("jokeHandler called from", r.RemoteAddr)
		joke := jokeService.GetJoke()

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(joke); err != nil {
			http.Error(w, "Failed to encode joke", http.StatusInternalServerError)
		}
	}
}
