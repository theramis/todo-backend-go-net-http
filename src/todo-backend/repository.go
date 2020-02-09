package main

import "errors"

var allTodos = make(map[int]*Todo)
var globalTodoId = 0

func addTodo(todo *Todo) {
	defer func() { globalTodoId++ }()

	todo.Id = globalTodoId
	allTodos[todo.Id] = todo
}

func getTodos() []Todo {
	result := make([]Todo, 0)
	for _, todo := range allTodos {
		result = append(result, *todo)
	}
	return result
}

func getTodo(id int) (Todo, error) {
	todo, ok := allTodos[id]
	if !ok {
		return Todo{}, errors.New("todo not found")
	}
	return *todo, nil
}

func deleteTodo(id int) {
	delete(allTodos, id)
}

func deleteAllTodos() {
	allTodos = make(map[int]*Todo)
}

func updateTodo(todo Todo) {
	allTodos[todo.Id] = &todo
}
