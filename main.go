package main

import (
	"net/http"

	"johannes-jahn.com/demo/controllers"
	"johannes-jahn.com/demo/joke"
	"johannes-jahn.com/demo/logger"
)

func main() {
	jokeSvc := joke.NewService()
	mux := http.NewServeMux()
	controllers.Register(mux, jokeSvc)

	logger.Println("Server starting on port 3000... 🚀")
	if err := http.ListenAndServe(":3000", mux); err != nil {
		logger.Fatalf("Could not start server: %s\n", err)
	}
}

