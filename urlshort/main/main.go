package main

import (
	"fmt"
	"net/http"

	"github.com/janhaans/learngo/urlshort"
)

func main() {
	mux := DefaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution`

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	yamlHandler := urlshort.YAMLHandler(yaml, mapHandler)
	http.ListenAndServe(":8080", yamlHandler)

}

func DefaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!!</h1>")
}
