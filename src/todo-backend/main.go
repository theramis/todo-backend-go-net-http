package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	addCorsHeaders(&writer)

	switch request.Method {
	case "POST":
		createTodo(&writer, request)
	default:
		_, _ = fmt.Fprint(writer, "Hello World!")
	}
}

func createTodo(writer *http.ResponseWriter, request *http.Request) {
	w := *writer
	w.Header().Set("Content-Type", "application/json")
	contentLength, _ := strconv.Atoi(request.Header.Get("Content-Length"))
	bodyBytes := make([]byte, contentLength)
	request.Body.Read(bodyBytes)
	w.Write(bodyBytes)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Env var 'PORT' must be set")
	}

	http.HandleFunc("/todo", handler)
	http.HandleFunc("/", catchAllHandler)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func catchAllHandler(writer http.ResponseWriter, request *http.Request) {
	addCorsHeaders(&writer)
	_, _ = fmt.Fprint(writer, "Hello World!")
}

func addCorsHeaders(writer *http.ResponseWriter) {
	w := *writer
	w.Header().Set("access-control-allow-origin", "*")
	w.Header().Set("access-control-allow-methods", "GET, POST, PATCH, DELETE")
	w.Header().Set("access-control-allow-headers", "accept, content-type")
}