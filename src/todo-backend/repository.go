package main

import "errors"

var allTodos = make(map[int]Todo)

func addTodo(todo Todo) error {
	_, ok := allTodos[todo.Id]
	if ok {
		return errors.New("todo already exists")
	}
	allTodos[todo.Id] = todo
	return nil
}

func getTodos() []Todo {
	result := make([]Todo, 0)
	for _, todo := range allTodos {
		result = append(result, todo)
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
