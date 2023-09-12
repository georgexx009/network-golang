package basenode

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func MakeHTTPRequest[T any](fullUrl string, httpMethod string, headers map[string]string, queryParameters url.Values, body io.Reader, responseType T) (T, error) {
	client := http.Client{}

	// check if the url is valid
	url, err := url.Parse(fullUrl)
	if err != nil {
		return responseType, err
	}

	// if the method is GET, lets append the query parameters
	if httpMethod == "GET" {
		q := url.Query()

		for k, v := range queryParameters {
			q.Set(k, strings.Join(v, ","))
		}

		url.RawQuery = q.Encode()
	}
}
