package app

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	mapURL(r)
	return r
}

//StartApp start the romannumber app
func StartApp() {
	router := setupRouter()
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
