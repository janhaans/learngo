//Package controllers have the controllers of example REST API application
package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/janhaans/learngo/mvc/utils"

	"github.com/janhaans/learngo/mvc/services"
)

//GetUser gets the user
func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 0, 64)
	if err != nil {
		//Return status code Bad Request
		userErr := &utils.ApplicationError{
			Message:    "user_id must be number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.Respond(c, userErr.StatusCode, userErr)
		return
	}
	fmt.Printf("About to process userid %d\n", userID)

	user, userErr := services.UsersService.GetUser(userID)
	if userErr != nil {
		//Return status code Not Found
		utils.Respond(c, userErr.StatusCode, userErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
