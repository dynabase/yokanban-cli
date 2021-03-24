package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"yokanban-cli/internal/auth"
	"yokanban-cli/internal/consts"
	"yokanban-cli/internal/http"
	"yokanban-cli/internal/utils"
)

// GetAccessToken retrieves either an access token from cache or creates a new one.
func GetAccessToken() string {
	log.Debug("getAccessToken")
	if cachedToken := getCachedAccessToken(); cachedToken != "" {
		log.Debug("\t getAccessToken - return cached access token")
		return cachedToken
	}
	return createNewAccessToken()
}

func getCachedAccessToken() string {
	log.Debug("getCachedAccessToken")
	exists, err := existsCachedAccessToken()
	if err != nil {
		log.Fatal(err)
	}
	if exists == false {
		log.Debug("\t getCachedAccessToken - cached access token does not exist")
		return ""
	}

	jsonFile, err := os.Open(getCachedAccessTokenFileURI())
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	var accessTokenData http.TokenData

	if err := json.Unmarshal(byteValue, &accessTokenData); err != nil {
		log.Fatal(err)
	}

	return accessTokenData.AccessToken
}

func createNewAccessToken() string {
	log.Debug("createNewAccessToken")
	jwt := auth.GetServiceAccountJWT()
	tokenData := http.Auth(jwt)

	// persist token to configuration directory for caching purposes
	tokenDataJSON, _ := json.Marshal(tokenData)
	if err := ioutil.WriteFile(getCachedAccessTokenFileURI(), tokenDataJSON, 0700); err != nil {
		log.Fatal(err)
	}

	return tokenData.AccessToken
}

func getCachedAccessTokenFileURI() string {
	log.Debug("getCachedAccessTokenFileURI")
	accessTokenFileURI := path.Join(utils.GetConfigDir(), consts.CachedTokenFilename)
	return accessTokenFileURI
}

func existsCachedAccessToken() (bool, error) {
	log.Debug("existsCachedAccessToken")
	accessTokenFileURI := getCachedAccessTokenFileURI()
	log.Debug("\t existsCachedAccessToken - check: " + accessTokenFileURI)

	stat, err := os.Stat(accessTokenFileURI)
	if err == nil {
		if stat != nil {
			log.Debug("\t\t existsCachedAccessToken - cached access token exists")
			return true, nil
		}
		return false, nil
	} else if os.IsNotExist(err) {
		log.Debug("\t\t existsCachedAccessToken - cached access token does not exist")
		return false, nil
	}

	// e.g. permission denied
	return false, err
}
