package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	var err error

	addCorsHeaders(&writer)
	switch request.Method {
	case "POST":
		err = createTodoHandler(&writer, request)
	case "GET":
		err = getTodosHandler(&writer, request)
	default:
		_, err = fmt.Fprint(writer, "Hello World!")
	}

	if err != nil {
		writer.WriteHeader(500)
		_, _ = fmt.Fprintf(writer, "Error processing request. %v", err)
	}
}

func createTodoHandler(writer *http.ResponseWriter, request *http.Request) error {
	w := *writer

	todo := Todo{}
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		return err
	}
	err = addTodo(todo)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todo)
	return err
}

func getTodosHandler(writer *http.ResponseWriter, request *http.Request) error {
	w := *writer
	todos := getTodos()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(todos)
	return err
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Env var 'PORT' must be set")
	}

	http.HandleFunc("/todos", handler)
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