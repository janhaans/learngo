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

//OptionsRomanNumber provide options of romannumber API
func OptionsRomanNumber(c *gin.Context) {
	c.Header("Allow", "GET, OPTIONS")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control_allow_headers", "origin, content-type, accept")
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)
}
