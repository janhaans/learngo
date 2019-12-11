package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//StartApp start the romannumber app
func StartApp() {
	router = gin.Default()
	mapURL()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
