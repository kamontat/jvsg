package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// MirrorBody will return the request as a response
func MirrorBody(w http.ResponseWriter, r *http.Request) (string, error) {
	// Throw if http method is not POST
	if err := ValidatePostMethod(w, r); err != nil {
		return "", err
	}

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	w.Header().Add("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Add("Content-Length", fmt.Sprintf("%d", len(resp)))
	_, err = w.Write(resp)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("response length: %d\n", len(resp)), nil
}

var cacheJson, _ = os.ReadFile("large.json")

// GetJsonBody will return large json file as a body
func GetCacheJsonBody(w http.ResponseWriter, r *http.Request) (string, error) {
	// Throw if http method is not POST
	if err := ValidatePostMethod(w, r); err != nil {
		return "", err
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Length", fmt.Sprintf("%d", len(cacheJson)))
	w.WriteHeader(http.StatusOK)
	w.Write(cacheJson)

	return fmt.Sprintf("response length: %d\n", len(cacheJson)), nil
}

// GetJsonBody will return large json file as a body
func GetJsonBody(w http.ResponseWriter, r *http.Request) (string, error) {
	// Throw if http method is not POST
	if err := ValidatePostMethod(w, r); err != nil {
		return "", err
	}

	json, err := os.ReadFile("large.json")
	if err != nil {
		return "", err
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Length", fmt.Sprintf("%d", len(json)))
	w.WriteHeader(http.StatusOK)
	w.Write(json)

	return fmt.Sprintf("response length: %d\n", len(json)), nil
}
