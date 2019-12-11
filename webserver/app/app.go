package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//StartApp start static webserver
func StartApp() {
	router = gin.Default()
	mapURL()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
