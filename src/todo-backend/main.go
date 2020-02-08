package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprint(writer, "Hello World!")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
