package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// RawRequest without parse response
func RawRequest(path, debug string) (io.ReadCloser, error) {
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

	if res.StatusCode != http.StatusOK {
		fmt.Printf("error response: %s\n", res.Status)
		return nil, errors.New("response status is not ok (" + res.Status + ")")
	}

	return res.Body, nil
}

func RequestWithoutParse(path, debug string) (string, error) {
	start := time.Now()
	body, err := RawRequest(path, debug)
	if err != nil {
		return "", err
	}

	result, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}

	RequestDurationMetric(start)
	return string(result), nil
}

// Request with parse response to map
func RequestWithParse(path, debug string) ([]map[string]interface{}, error) {
	start := time.Now()
	result := []map[string]interface{}{}
	body, err := RawRequest(path, debug)
	if err != nil {
		return result, err
	}

	json.NewDecoder(body).Decode(&result)

	RequestWithJsonDurationMetric(start)
	return result, nil
}
