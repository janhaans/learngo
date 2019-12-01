//Package controllers have the controllers of example REST API application
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/janhaans/learngo/mvc/utils"

	"github.com/janhaans/learngo/mvc/services"
)

//GetUser gets the user
func GetUser(res http.ResponseWriter, req *http.Request) {
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 0, 64)
	if err != nil {
		//Return status code Bad Request
		userErr := &utils.ApplicationError{
			Message:    "user_id must be number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		res.WriteHeader(userErr.StatusCode)
		jsonValue, _ := json.Marshal(userErr)
		res.Write(jsonValue)
		return
	}
	fmt.Printf("About to process userid %d\n", userID)

	user, userErr := services.UsersService.GetUser(userID)
	if userErr != nil {
		//Return stataus code Not Found
		res.WriteHeader(userErr.StatusCode)
		jsonValue, _ := json.Marshal(userErr)
		res.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)

}
