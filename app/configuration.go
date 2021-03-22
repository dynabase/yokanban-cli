package app

import (
	"errors"
	"os"
)

func GetApiUrl() string {
	if value, ok := os.LookupEnv("YOKANBAN_API_URL"); ok {
		return value
	}
	return "https://api.yokanban.io"
}

func GetApiKeysPath() (string, error) {
	if value, ok := os.LookupEnv("YOKANBAN_API_KEYS_PATH"); ok {
		return value, nil
	}
	return "", errors.New("env var YOKANBAN_API_KEYS_PATH not defined")
}
