package main

import (
	"net/http"
	"os"
)

var debugMode = GetEnv("DEBUG", "false")

func GetEnv(key string, def string) string {
	val, exist := os.LookupEnv(key)
	if !exist {
		return def
	}

	return val
}

func IsDebug(r *http.Request) bool {
	if debugMode == "true" {
		return true
	}

	var query = r.URL.Query()
	return query.Has("debug") && query.Get("debug") == "true"
}
