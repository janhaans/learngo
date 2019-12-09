package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enabledMocks = false
	mocks        = make(map[string]*Mock)
)

type Mock struct {
	Url        string
	HttpMethod string
	Response   *http.Response
	Err        error
}

//StartMockups mocks Post
func StartMockups() {
	enabledMocks = true
}

//StopMockups stops mocking Post
func StopMockups() {
	enabledMocks = false
}

func AddMock(mock Mock) {
	mocks[getMockID(mock.HttpMethod, mock.Url)] = &mock
}

func getMockID(httpMethod string, url string) string {
	return fmt.Sprintf("%s-%s", httpMethod, url)
}

//Post makes HTTP POST request
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		mock := mocks[getMockID(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("No mock found for url")
		}
		return mock.Response, mock.Err
	}
	bodybytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodybytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
