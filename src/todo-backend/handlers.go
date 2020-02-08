package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var globalId = 0

func createTodoHandler(writer *http.ResponseWriter, request *http.Request) error {
	w := *writer

	todo := Todo{}
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		return err
	}
	err = addTodo(globalId, todo)
	if err != nil {
		return err
	}
	globalId++

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todo)
	return err
}

func getTodosHandler(writer *http.ResponseWriter, request *http.Request) error {
	w := *writer
	todos := getTodos()

	result := make([]TodoResponse, 0)
	for key, todo := range todos {
		url := fmt.Sprintf("http://%v/todos/%v", request.Host, strconv.Itoa(key))
		result = append(result, TodoResponse{Todo: todo, Url: url})
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(result)
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
