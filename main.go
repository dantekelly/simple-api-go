package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	initializeServer()
}

func initializeServer() {
	fmt.Println("Starting server on port :3333")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		helloWorld(w)
	})

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func writeMessage(w http.ResponseWriter, message string) {
	_, err := w.Write([]byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func helloWorld(w http.ResponseWriter) {
	// String to return
	s := "Hello, World!"

	writeMessage(w, s)
	fmt.Println(s)
}
