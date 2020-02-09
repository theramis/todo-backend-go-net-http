package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	var err error
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
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Env var 'PORT' must be set")
	}

	mux := http.NewServeMux()

	mux.Handle("/todos", addAllMiddlewares(handler))
	mux.Handle("/todos/", addAllMiddlewares(handler))
	//mux.HandleFunc("/", catchAllHandler)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
