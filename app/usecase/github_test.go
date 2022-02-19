package usecase

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Dionizio8/example-mocking-http-request/infra/restclient"
	"github.com/Dionizio8/example-mocking-http-request/infra/restclient/mocks"
	"github.com/stretchr/testify/assert"
)

func init() {
	restclient.Client = &mocks.MockClient{}
}
func TestGetReposSuccess(t *testing.T) {
	json := `[{"id": 88888888, "name": "Test Name Repo","private": false}]`
	response := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       response,
		}, nil
	}

	resp, err := GetRepos("Test Name Repo")

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Equal(t, "Test Name Repo", resp[0]["name"])
	assert.Equal(t, float64(88888888), resp[0]["id"])
}

func TestGetReposError(t *testing.T) {
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New("repo not found")
	}

	_, err := GetRepos("Test Name Repo")

	assert.NotNil(t, err)
}
