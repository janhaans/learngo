package app

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/janhaans/learngo/romannumber/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetRomanNumber(t *testing.T) {
	router := setupRouter()

	req := httptest.NewRequest("GET", "/romannumber/4", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//resp := w.Result()
	assert.Equal(t, 200, w.Code)
	var romannumber domain.RomanNumber
	err := json.Unmarshal(w.Body.Bytes(), &romannumber)
	assert.Nil(t, err)
	assert.EqualValues(t, 4, romannumber.IntNumber)
	assert.EqualValues(t, "IV", romannumber.RomanNumber)

	/*	Alternative way to test:
		resp := w.Result()
		assert.Equal(t, 200, resp.StatusCode)
		body, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		assert.EqualValues(t, `{"integer":4,"roman-number":"IV"}`, strings.TrimSpace(string(body)))
	*/
}
