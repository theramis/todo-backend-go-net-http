package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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

func deleteTodoHandler(rawId string) error {
	if rawId == "" {
		deleteAllTodos()
		return nil
	}
	id, err := strconv.Atoi(rawId)
	if err != nil {
		return err
	}
	deleteTodo(id)
	return nil
}

func catchAllHandler(writer http.ResponseWriter, request *http.Request) {
	addCorsHeaders(&writer)

	for _, v := range allRequests {
		_, _ = fmt.Fprintln(writer, v)
	}
}

func addCorsHeaders(writer *http.ResponseWriter) {
	w := *writer
	w.Header().Set("access-control-allow-origin", "*")
	w.Header().Set("access-control-allow-methods", "GET, POST, PATCH, DELETE")
	w.Header().Set("access-control-allow-headers", "accept, content-type")
}
