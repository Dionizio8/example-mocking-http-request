package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Dionizio8/example-mocking-http-request/infra/restclient"
)

func GetRepos(username string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos?sort=created&direction=desc", username)
	response, err := restclient.Request(http.MethodGet, url, nil, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("repo not found")
	}

	defer response.Body.Close()
	m := []map[string]interface{}{}
	err = json.NewDecoder(response.Body).Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
