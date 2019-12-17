package domain

import (
	"net/http"

	"github.com/janhaans/learngo/romannumber/utils"
)

/*
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
*/

func intToRoman(num int) string {
	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I"}
	roman := ""
	i := 0

	for num > 0 {
		// calculate the number of times this num is completly divisible by values[i]
		// times will only be > 0, when num >= values[i]
		k := num / values[i]
		for j := 0; j < k; j++ {
			//buildup roman numeral
			roman += symbols[i]

			//reduce the value of num.
			num -= values[i]
		}
		i++
	}
	return roman
}

//GetRomanNumber returns Roman Number or error when roman number is not found
func GetRomanNumber(num int64) (*RomanNumber, *utils.ApplicationError) {
	roman := intToRoman(int(num))
	if roman == "" {
		return nil, &utils.ApplicationError{
			Code:        "Error1",
			StatusCode:  http.StatusBadRequest,
			Description: "Number is less then 1",
		}
	}
	return &RomanNumber{
		IntNumber:   num,
		RomanNumber: roman,
	}, nil

	/*
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
	*/
}
