package api

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

const (
	caasBaseURL     = "https://cataas.com/"
	caasCatEndpoint = "cat"
)

func RequestCat(timeout time.Duration) ([]byte, error) {
	// make some stuff
	bodyReader := bytes.NewReader(make([]byte, 0))
	reqURL := caasBaseURL + caasCatEndpoint
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest(http.MethodGet, reqURL, bodyReader)
	if err != nil {
		return nil, err
	}
	// do the thing
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// clean up
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {

		}
	}(resp.Body)

	// hey it's a cat!
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
