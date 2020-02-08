package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("access-control-allow-origin", "*")
	writer.Header().Set("access-control-allow-methods", "GET, POST, PATCH, DELETE")
	writer.Header().Set("access-control-allow-headers", "accept, content-type")
	_, _ = fmt.Fprint(writer, "Hello World!")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Env var 'PORT' must be set")
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
