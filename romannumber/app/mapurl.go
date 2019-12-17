package app

import (
	"github.com/gin-gonic/gin"
	"github.com/janhaans/learngo/romannumber/controllers"
)

func mapURL(router *gin.Engine) {
	router.GET("/romannumber/:num", controllers.GetRomanNumber)
	router.OPTIONS("romannumber", controllers.OptionsRomanNumber)
}
