package accesstoken

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"yokanban-cli/internal/auth"
	"yokanban-cli/internal/consts"
	yohttp "yokanban-cli/internal/http"
	"yokanban-cli/internal/utils"

	log "github.com/sirupsen/logrus"
)

// Get retrieves either an access token from cache or creates a new one.
func Get() string {
	log.Debug("Get")
	if cachedToken := getCachedAccessToken(); cachedToken != "" {
		log.Debug("\t Get - return cached access token")
		return cachedToken
	}
	return Refresh()
}

// Refresh creates a new access token and overwrites cached one.
func Refresh() string {
	log.Debug("Refresh")
	jwt := auth.GetServiceAccountJWT()
	h := yohttp.HTTP{Client: &http.Client{}}
	tokenData, err := h.Auth(jwt)
	if err != nil {
		log.Fatal(err)
	}

	// persist token to configuration directory for caching purposes
	tokenDataJSON, _ := json.Marshal(tokenData)
	if err := ioutil.WriteFile(getCachedAccessTokenFileURI(), tokenDataJSON, 0700); err != nil {
		log.Fatal(err)
	}

	return tokenData.AccessToken
}

func getCachedAccessToken() string {
	log.Debug("getCachedAccessToken")
	exists, err := existsCachedAccessToken()
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
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

	var accessTokenData yohttp.TokenData

	if err := json.Unmarshal(byteValue, &accessTokenData); err != nil {
		log.Fatal(err)
	}

	return accessTokenData.AccessToken
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
