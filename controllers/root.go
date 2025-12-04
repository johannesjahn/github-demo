package controllers

import (
	"net/http"

	"johannes-jahn.com/demo/joke"
	"johannes-jahn.com/demo/logger"
)

func rootHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		logger.Println("rootHandler called from", r.RemoteAddr)
		w.Write([]byte("Hello!"))
	}
}

func Register(mux *http.ServeMux, jokeService joke.Service) {
	mux.HandleFunc("/", rootHandler())
	mux.HandleFunc("/joke", jokeHandler(jokeService))
}
