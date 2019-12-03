package utils

import (
	"github.com/gin-gonic/gin"
)

//Respond gives XML or JSON response based on Accept header
func Respond(c *gin.Context, statuscode int, body interface{}) {
	if c.GetHeader("Accept") == "application/xml" {
		c.XML(statuscode, body)
		return
	}
	c.JSON(statuscode, body)
	return
}
