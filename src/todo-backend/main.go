package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Env var 'PORT' must be set")
	}

	mux := http.NewServeMux()
	mux.Handle("/todos", addAllMiddlewares(rootHandler))
	mux.Handle("/todos/", addAllMiddlewares(resourceHandler))

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
