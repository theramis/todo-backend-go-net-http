package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		createTodoHandler(w, r)
	case "GET":
		getAllTodosHandler(w, r)
	case "DELETE":
		deleteAllTodosHandler()
	default:
		http.NotFoundHandler().ServeHTTP(w, r)
	}
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	if path := strings.Split(r.URL.Path[1:], "/"); len(path) == 2 {
		rawId := path[1]
		id, err := strconv.Atoi(rawId)
		if err != nil {
			w.WriteHeader(400)
			_, _ = w.Write([]byte(fmt.Sprintf("Invalid todo id given: '%v'", rawId)))
			return
		}

		switch r.Method {
		case "GET":
			getTodoHandler(w, r, id)
		case "PATCH":
			updateTodoHandler(w, r, id)
		case "DELETE":
			deleteTodoHandler(id)
		default:
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	} else {
		panic(errors.New("reached resource handler but shouldn't have"))
	}
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	todo := Todo{}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		errorResponse(w, err)
		return
	}

	addTodo(&todo)
	todo.setUrl(r)

	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		errorResponse(w, err)
		return
	}
}

func getAllTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos := getTodos()
	for i, todo := range todos {
		todos[i] = todo
	}
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		errorResponse(w, err)
	}
}

func getTodoHandler(w http.ResponseWriter, r *http.Request, id int) {
	todo, err := getTodo(id)
	if err != nil {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}
	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		errorResponse(w, err)
	}
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request, id int) {
	todo, err := getTodo(id)
	if err != nil {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}

	updatedTodo := Todo{}

	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		errorResponse(w, err)
		return
	}

	todo.Order = updatedTodo.Order
	todo.Completed = updatedTodo.Completed
	todo.Title = updatedTodo.Title

	updateTodo(todo)

	err = json.NewEncoder(w).Encode(todo)
	if err != nil {
		errorResponse(w, err)
	}
}

func deleteTodoHandler(id int) {
	deleteTodo(id)
}

func deleteAllTodosHandler() {
	deleteAllTodos()
}

func errorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}
