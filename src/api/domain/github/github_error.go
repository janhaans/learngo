package github

//ErrorResponse models Github error response
type ErrorResponse struct {
	StatusCode       int     `json:"status_code"`
	Message          string  `json:"message"`
	DocumentationURL string  `json:"documentation_url"`
	Errors           []Error `json:"errors"`
}

//Error models Github error
type Error struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}

/*
{
    "message": "Repository creation failed.",
    "errors": [
        {
            "resource": "Repository",
            "code": "custom",
            "field": "name",
            "message": "name already exists on this account"
        }
    ],
    "documentation_url": "https://developer.github.com/v3/repos/#create"
}
*/
