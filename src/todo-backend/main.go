package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var allRequests = make([]string, 0)

func handler(writer http.ResponseWriter, request *http.Request) {
	allRequests = append(allRequests, request.Method+": "+request.URL.Path+"|"+request.RequestURI)

	var err error
	addCorsHeaders(&writer)

	todoId := ""
	if path := strings.Split(request.URL.Path[1:], "/"); len(path) == 2 {
		todoId = path[1]
	}

	switch request.Method {
	case "POST":
		err = createTodoHandler(&writer, request)
	case "GET":
		err = getTodosHandler(&writer, request)
	case "DELETE":
		deleteTodoHandler(todoId)
	default:
		_, err = fmt.Fprint(writer, "You are not mapped yet!")
	}

	if err != nil {
		writer.WriteHeader(500)
		_, _ = fmt.Fprintf(writer, "Error processing request. %v", err)
	}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Env var 'PORT' must be set")
	}

	http.HandleFunc("/todos", handler)
	http.HandleFunc("/todos/", handler)
	http.HandleFunc("/", catchAllHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
