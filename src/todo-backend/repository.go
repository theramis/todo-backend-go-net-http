package main

import "errors"

var allTodos = make(map[int]Todo)

func addTodo(id int, todo Todo) error {
	_, ok := allTodos[id]
	if ok {
		return errors.New("todo already exists")
	}
	allTodos[id] = todo
	return nil
}

func getTodos() map[int]Todo {
	result := make([]Todo, 0)
	for _, todo := range allTodos {
		result = append(result, todo)
	}
	return allTodos
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
