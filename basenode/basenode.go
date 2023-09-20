package basenode

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func MakeHTTPRequest[T any](fullUrl string, httpMethod string, headers map[string]string, queryParameters url.Values, body io.Reader, responseType T) (T, error) {
	client := http.Client{}

	// check if the url is valid
	u, err := url.Parse(fullUrl)
	if err != nil {
		return responseType, err
	}

	// if the method is GET, lets append the query parameters
	if httpMethod == "GET" {
		q := u.Query()

		for k, v := range queryParameters {
			q.Set(k, strings.Join(v, ",")) // the join is for arrays
		}

		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(httpMethod, u.String(), body)
	if err != nil {
		return responseType, err
	}

	// add headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// do the request
	res, err := client.Do(req)
	if err != nil {
		return responseType, err
	}
	if res == nil {
		return responseType, fmt.Errorf("error: calling %s returned empty response", u.String())
	}

	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return responseType, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return responseType, fmt.Errorf("error: calling %s\nstatus: %s\nresponseData: %s", u.String(), res.StatusCode, responseData)
	}

	var responseObject T
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Printf("error unmarshalling response: %+v", err)
		return responseType, err
	}

	return responseObject, nil
}
