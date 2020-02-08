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

	todo := Todo{Id: globalId}
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		return err
	}
	err = addTodo(todo)
	defer func() { globalId++ }()
	if err != nil {
		return err
	}

	todo.setUrl(request)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(todo)
	return err
}

func getTodosHandler(writer *http.ResponseWriter, request *http.Request, rawId string) error {
	w := *writer
	var err error
	w.Header().Set("Content-Type", "application/json")
	if rawId == "" {
		todos := getTodos()

		for i, todo := range todos {
			todo.setUrl(request)
			todos[i] = todo
		}

		err = json.NewEncoder(w).Encode(todos)
	} else {
		id, err := strconv.Atoi(rawId)
		if err != nil {
			return err
		}

		todo, err := getTodo(id)
		if err != nil {
			return err
		}
		todo.setUrl(request)
		err = json.NewEncoder(w).Encode(todo)
	}

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
