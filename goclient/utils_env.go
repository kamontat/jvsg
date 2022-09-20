package main

import (
	"os"
)

func GetEnv(key string, def string) string {
	val, exist := os.LookupEnv(key)
	if !exist {
		return def
	}

	return val
}
