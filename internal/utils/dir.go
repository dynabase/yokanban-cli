package utils

import (
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

/**
Retrieves the OS specific home directory of the current user.
*/
func GetUserHomeDir() string {
	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

/**
Checks whether a directory exists or not.
*/
func ExistsDir(dir string) (bool, error) {
	stat, err := os.Stat(dir)
	if err == nil {
		if stat != nil {
			return true, nil
		}
		return false, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	// e.g. permission denied
	return false, err
}

/**
Retrieves the yokanban configuration directory. It will be created if it doesn't exist.
*/
func GetConfigDir() string {
	homeDir := GetUserHomeDir()
	configDir := path.Join(homeDir, ".yokanban")

	exists, err := ExistsDir(configDir)
	if err != nil {
		log.Fatal(err)
	}
	if exists == false {
		if err := os.MkdirAll(configDir, 0700); err != nil {
			log.Fatal(err)
		}
	}

	return configDir
}
