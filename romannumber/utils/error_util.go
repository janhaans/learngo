package utils

//ApplicationError models application error
type ApplicationError struct {
	Code        string `json:"code"`
	StatusCode  int    `json:"status"`
	Description string `json:"description"`
}
