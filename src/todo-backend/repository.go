package main

import "errors"

var allTodos = make(map[int]Todo)

func addTodo(todo Todo) error {
	_, ok := allTodos[todo.Order]
	if ok {
		return errors.New("todo already exists")
	}
	allTodos[todo.Order] = todo
	return nil
}

func getTodos() []Todo {
	result := make([]Todo, len(allTodos))
	for i, todo := range allTodos {
		result[i-1] = todo
	}
	return result
}

func getTodo(id int) (Todo, error) {
	todo, ok := allTodos[id]
	if !ok {
		return Todo{}, errors.New("todo not found")
	}
	return todo, nil
}

func deleteTodo(id int) {
	delete(allTodos, id)
}

func deleteAllTodos() {
	allTodos = make(map[int]Todo)
}
