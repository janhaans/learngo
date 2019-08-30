package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "<h1>Hello, Hello, Hello</h1>")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "<h1>World, World, World</h1>")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	hello := HelloHandler{}
	world := WorldHandler{}

	http.Handle("/hello", &hello)
	http.Handle("/world/", &world)

	server.ListenAndServe()
}
