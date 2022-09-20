package main

import (
	"errors"
	"net/http"
)

func ValidatePostMethod(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return errors.New("only Post method is allow")
	}

	return nil
}
