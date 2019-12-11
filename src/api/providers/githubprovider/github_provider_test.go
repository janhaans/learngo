package githubprovider

import (
	"errors"
	"net/http"
	"testing"

	"github.com/janhaans/learngo/src/api/client/restclient"
	"github.com/janhaans/learngo/src/api/domain/github"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizationHeaderValue(t *testing.T) {
	header := getAuthorizationHeaderValue("123")
	assert.EqualValues(t, "token 123", header)
}

func TestCreateRepoErrorRestClient(t *testing.T) {
	restclient.StartMockups()
	restclient.AddMock(restclient.Mock{
		Url:        "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid restclient response", err.Message)
}

func TestCreateRepoInvalidBody(t *testing.T) {

}
