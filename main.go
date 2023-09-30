package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Server struct {
	Router *chi.Mux
}

func main() {
	s := CreateNewServer()
	s.MountHandlers()

	err := http.ListenAndServe(":3333", s.Router)

	if err != nil {
		fmt.Println(err)

		return
	}
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()

	return s
}

func (s *Server) MountHandlers() {
	s.Router.Use(middleware.Logger)

	s.Router.Get("/hello-world", HelloWorld)
}

func writeMessage(w http.ResponseWriter, message string) {
	_, err := w.Write([]byte(message))

	if err != nil {
		fmt.Println(err)
		return
	}
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// String to return
	s := "Hello, World!"

	writeMessage(w, s)
	// fmt.Println(s)
}
