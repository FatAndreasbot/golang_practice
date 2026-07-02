package server

import (
	"fmt"
	"net/http"
)

type Handler struct {
	path    string
	handler func(http.ResponseWriter, *http.Request)
}

type Server []Handler

func (s *Server) Init() {
	for _, handler := range *s {
		http.HandleFunc(handler.path, handler.handler)
	}
}

func (s *Server) Launch(port int) {
	fmt.Printf("launching server at %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
