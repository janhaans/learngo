package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRomanNumber(t *testing.T) {
	roman, err := GetRomanNumber(1)
	assert.Nil(t, err)
	assert.NotNil(t, roman)
	assert.EqualValues(t, "I", roman.RomanNumber)
}

func TestGetRomanNumberError(t *testing.T) {
	roman, err := GetRomanNumber(0)
	assert.Nil(t, roman)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Error1", err.Code)
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode)
	assert.EqualValues(t, "Number is less then 1", err.Description)
}
