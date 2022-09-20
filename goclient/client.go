package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Request without parse response
func Request(path, debug string) (io.ReadCloser, error) {
	start := time.Now()
	url := GetUrl(path, debug)

	req, err := GetRequest(url)
	if err != nil {
		fmt.Printf("error creating http request: %s\n", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return nil, err
	}

	RequestDurationMetric(start)
	if res.StatusCode != http.StatusOK {
		fmt.Printf("error response: %s\n", res.Status)
		return nil, errors.New("response status is not ok (" + res.Status + ")")
	}

	return res.Body, nil
}

// Request with parse response to map
func RequestWithParse(path, debug string) ([]map[string]interface{}, error) {
	start := time.Now()
	result := []map[string]interface{}{}
	body, err := Request(path, debug)
	if err != nil {
		return result, err
	}
	json.NewDecoder(body).Decode(&result)
	RequestWithJsonDurationMetric(start)

	return result, nil
}
