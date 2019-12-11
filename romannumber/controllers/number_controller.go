package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/janhaans/learngo/romannumber/domain"
	"github.com/janhaans/learngo/romannumber/utils"
)

//GetRomanNumber get roman number
func GetRomanNumber(c *gin.Context) {
	num, err := strconv.ParseInt(c.Param("num"), 0, 64)
	if err != nil {
		inputErr := &utils.ApplicationError{
			Code:        "bad request",
			StatusCode:  http.StatusBadRequest,
			Description: "Must be number",
		}
		c.JSON(inputErr.StatusCode, inputErr)
		return
	}

	romanNumber, applicationError := domain.GetRomanNumber(num)
	if applicationError != nil {
		c.JSON(applicationError.StatusCode, applicationError)
		return
	}
	c.JSON(http.StatusOK, romanNumber)
	return
}
