package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewServer()
	http.HandleFunc("/chat", server.Handler)

	log.Println("Server started at localhost:8000")
	log.Fatal(
		http.ListenAndServe("localhost:8000", nil),
	)
}
