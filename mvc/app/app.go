//Package app is example REST API application
package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

//StartApp start the REST API application
func StartApp() {
	mapURLs()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
