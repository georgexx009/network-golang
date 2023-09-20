package main

import (
	"fmt"
	"net/url"
	"network-golang/basenode"
)

func main() {
	rawUrl := "https://api.github.com/users/georgexx009"
	headers := map[string]string{
		"Accept": "application/vnd.github.v3+json",
	}
	queryParams := url.Values{}
	queryParams.Add("per_page", "1")

	var response map[string]interface{}

	response, err := basenode.MakeHTTPRequest(rawUrl, "GET", headers, queryParams, nil, response)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response: %+v", response)
}
