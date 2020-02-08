package main

type Todo struct {
	Title     string `json:"title"`
	Order     int    `json:"order"`
	Completed bool   `json:"completed"`
}
