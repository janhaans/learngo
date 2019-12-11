package domain

//RomanNumber models roman number
type RomanNumber struct {
	IntNumber   int64  `json:"integer"`
	RomanNumber string `json:"roman-number"`
}
