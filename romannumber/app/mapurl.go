package app

import "github.com/janhaans/learngo/romannumber/controllers"

func mapURL() {
	router.GET("/romannumber/:num", controllers.GetRomanNumber)
}
