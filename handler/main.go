package main

import (
	"fmt"
	"net/http"
)

//Handler is an interface with method ServeHTTP
//HellaHandler is a struct that implements the Handler interface
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

//Byehandler is a struct that implements the Handler interface
type ByeHandler struct{}

func (h *ByeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bye World")
}

func main() {

	//Configre the HTTP Server
	//No Handler will be specified for HTTP Server; default Handler of HTTP Server is DefaultServeMux
	server := http.Server{
		Addr: ":8080", //HTTP Server listens on any IP address at port 8080
	}

	//Create instance of type HelloHandler
	hello := HelloHandler{}
	//Create instance of type ByeHandler
	bye := ByeHandler{}

	//Add Handler hello to DefaultServeMux
	http.Handle("/hello", &hello)
	//Add Handler bye to DefaultServeMux
	http.Handle("/bye", &bye)

	//Start the HTTP Server
	server.ListenAndServe()
}
