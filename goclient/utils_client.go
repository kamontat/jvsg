package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var SERVER_HOST = GetEnv("SERVER_HOST", "localhost")
var SERVER_PORT = GetEnv("SERVER_PORT", "3333")

func GetUrl(path string, debug string) url.URL {
	return url.URL{
		Scheme:   "http",
		Host:     fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT),
		Path:     path,
		RawQuery: "debug=" + debug,
	}
}

func GetRequest(link url.URL) (*http.Request, error) {
	return GetRequestBody(link, http.NoBody)
}

func GetRequestBody(link url.URL, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest("POST", link.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	return req, nil
}
