package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template
var char = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"@", "#", "$", "&"}

func generatePassword() string {
	log.Println("Start password generator")
	passwd := ""
	for i := 0; i < 8; i++ {
		passwd = passwd + char[rand.Intn(65)]
	}
	log.Printf("Generated password = %s\n", passwd)
	log.Println("End password generator")
	return passwd
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	handleError(w, err)
}

func passwd(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "passwd.gohtml", generatePassword())
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Fprintf(w, "Error: %s - %d\n", err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	rand.Seed(time.Now().UnixNano())
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/passwd", passwd)
	router.ServeFiles("/css/*filepath", http.Dir("/home/stellar/go/src/github.com/janhaans/learngo/passwordgeneratorWebApp/css"))

	http.ListenAndServe(":8080", router)
}
