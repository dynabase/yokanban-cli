package config

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

// GetAPIURL retrieves the url of the yokanban HTTP API.
func GetAPIURL() string {
	if value, ok := os.LookupEnv("YOKANBAN_API_URL"); ok {
		return value
	}
	return "https://api.yokanban.io"
}

//GetAPIKeysPath retrieves the path to the API keys of a service account.
func GetAPIKeysPath() (string, error) {
	if value, ok := os.LookupEnv("YOKANBAN_API_KEYS_PATH"); ok {
		return value, nil
	}
	return "", errors.New("env var YOKANBAN_API_KEYS_PATH not defined")
}

// GetLogLevel retrieves the application log level.
// e.g. trace, debug, info, warn, error, fatal, panic
func GetLogLevel() logrus.Level {
	if value, ok := os.LookupEnv("YOKANBAN_LOGLEVEL"); ok {
		level, err := logrus.ParseLevel(value)
		if err != nil {
			fmt.Println(value + " is not a valid loglevel. Using default.")
			return GetDefaultLogLevel()
		}
		return level
	}
	return GetDefaultLogLevel()
}

// GetDefaultLogLevel retrieves the default application loglevel.
func GetDefaultLogLevel() logrus.Level {
	return logrus.WarnLevel
}
