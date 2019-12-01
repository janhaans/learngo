package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/janhaans/learngo/mvc/utils"
)

type usersDao struct{}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "Jan", LastName: "Haans", Email: "jan.haans@example.com"},
	}
	//UsersDao provides access to methods of interface usersDaointerface
	UsersDao usersDaoInterface
)

func init() {
	UsersDao = &usersDao{}
}

//GetUser gets user from database
func (u *usersDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("Query the database")
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user with ID %d was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
