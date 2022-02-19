package main

import (
	"fmt"
	"os"

	"github.com/Dionizio8/example-mocking-http-request/app/usecase"
)

func main() {
	repoName := os.Args[1]

	resp, err := usecase.GetRepos(repoName)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	fmt.Println(resp[0]["name"])
}
