package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := UsersDao.GetUser(0)

	assert.Nil(t, user, "Expected user to be nil")
	isNotNil := assert.NotNil(t, err, "Expected error to be not nil")
	if isNotNil {
		assert.EqualValues(t, "user with ID 0 was not found", err.Message)
		assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
		assert.EqualValues(t, "not_found", err.Code)
	}
}

func TestGetUserFound(t *testing.T) {
	user, err := UsersDao.GetUser(123)

	isNotNil := assert.NotNil(t, user, "Expected user to be nil")
	assert.Nil(t, err, "Expected error to be not nil")
	if isNotNil {
		assert.EqualValues(t, uint64(123), user.ID)
		assert.EqualValues(t, "Jan", user.FirstName)
		assert.EqualValues(t, "Haans", user.LastName)
		assert.EqualValues(t, "jan.haans@example.com", user.Email)
	}
}
