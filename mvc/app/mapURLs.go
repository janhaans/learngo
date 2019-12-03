package app

import "github.com/janhaans/learngo/mvc/controllers"

func mapURLs() {
	router.GET("/users/:user_id", controllers.GetUser)
}
