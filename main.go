package main

import (
	"log"
	"net/http"

	"johannes-jahn.com/demo/controllers"
	"johannes-jahn.com/demo/joke"
)

func main() {
	jokeSvc := joke.NewService()
	mux := http.NewServeMux()
	controllers.Register(mux, jokeSvc)

	log.Println("Server starting on port 3000...")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

