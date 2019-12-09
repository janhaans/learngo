package githubprovider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/janhaans/learngo/src/api/client/restclient"
	"github.com/janhaans/learngo/src/api/domain/github"
)

const (
	authorizationHeaderFormat = "token %s"
	createRepoURL             = "https://api.github.com/user/repos"
)

func getAuthorizationHeaderValue(accessToken string) string {
	return fmt.Sprintf(authorizationHeaderFormat, accessToken)
}

//CreateRepo creates Github repository
func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	headers := http.Header{}
	headers.Set("Authorization", getAuthorizationHeaderValue(accessToken))

	response, err := restclient.Post(createRepoURL, request, headers)
	fmt.Println(response)
	fmt.Println(err)
	if err != nil {
		log.Println(fmt.Sprintf("Error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid response body",
		}
	}
	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.ErrorResponse
		err = json.Unmarshal(bytes, &errResponse)
		if err != nil {
			return nil, &github.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		log.Println(fmt.Sprintf("Error when unmarshal create repo succesfull response %s", err.Error()))
		return nil, &github.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	return &result, nil
}
