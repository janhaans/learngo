//Package app is example REST API application
package app

import (
	"net/http"
	"github.com/janhaans/learngo/mvc/controllers"
)

//StartApp start the REST API application
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
