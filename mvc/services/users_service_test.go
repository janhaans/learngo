package services

import (
	"net/http"
	"testing"

	"github.com/janhaans/learngo/mvc/domain"
	"github.com/janhaans/learngo/mvc/utils"

	"github.com/stretchr/testify/assert"
)

var (
	//userDaoMock     usersDaoMock
	getUserFunction func(userID int64) (*domain.User, *utils.ApplicationError)
)

type usersDaoMock struct{}

func (u *usersDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

func init() {
	domain.UsersDao = &usersDaoMock{}
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userID int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 does not exist",
			StatusCode: http.StatusNotFound,
		}
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user, "Expected user to be nil")
	isNotNil := assert.NotNil(t, err, "Expected error to be not nil")
	if isNotNil {
		assert.EqualValues(t, "user 0 does not exist", err.Message)
		assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
		assert.EqualValues(t, "", err.Code)
	}
}
