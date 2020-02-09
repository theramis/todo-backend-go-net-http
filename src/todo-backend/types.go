package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type Todo struct {
	Title     string `json:"title"`
	Order     int    `json:"order"`
	Completed bool   `json:"completed"`
	Id        int    `json:"id"`
	Url       string `json:"url"`
}

func (todo *Todo) setUrl(request *http.Request) {
	todo.Url = fmt.Sprintf("https://%v/todos/%v", request.Host, strconv.Itoa(todo.Id))
}
