package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var allRequests = make([]string, 0)

func addRequest(request *http.Request) {
	err := recover()
	if err != nil {
		allRequests = append(allRequests, request.Method+": "+request.URL.Path+" Panic!")
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	defer addRequest(request)

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
		err = getTodosHandler(&writer, request, todoId)
	case "PATCH":
		err = updateTodoHandler(&writer, request, todoId)
	case "DELETE":
		err = deleteTodoHandler(todoId)
	default:
		_, err = fmt.Fprint(writer, "You are not mapped yet!")
	}

	if err != nil {
		writer.WriteHeader(500)
		_, _ = fmt.Fprintf(writer, "Error processing request. %v", err)
		allRequests = append(allRequests, request.Method+": "+request.URL.Path+" 500 - "+err.Error())
	} else {
		allRequests = append(allRequests, request.Method+": "+request.URL.Path+" 200")
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
