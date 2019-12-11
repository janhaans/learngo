package domain

import (
	"net/http"

	"github.com/janhaans/learngo/romannumber/utils"
)

var (
	romanNumbers = map[int64]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}
)

//GetRomanNumber returns Roman Number or error when roman number is not found
func GetRomanNumber(num int64) (*RomanNumber, *utils.ApplicationError) {
	roman, ok := romanNumbers[num]
	if ok {
		return &RomanNumber{
			IntNumber:   num,
			RomanNumber: roman,
		}, nil
	}
	return nil, &utils.ApplicationError{
		Code:        "Error1",
		StatusCode:  http.StatusBadRequest,
		Description: "Number is outside 1-10 range",
	}
}
