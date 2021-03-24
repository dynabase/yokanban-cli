package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"yokanban-cli/internal/auth"
	"yokanban-cli/internal/http"
	"yokanban-cli/internal/utils"
)

/**
Retrieves either an access token from cache or creates a new one.
*/
func GetAccessToken() string {
	fmt.Println("getAccessToken")
	if cachedToken := getCachedAccessToken(); cachedToken != "" {
		fmt.Println("\t getAccessToken - return cached access token")
		return cachedToken
	}
	return createNewAccessToken()
}

func getCachedAccessToken() string {
	fmt.Println("getCachedAccessToken")
	exists, err := existsCachedAccessToken()
	if err != nil {
		log.Fatal(err)
	}
	if exists == false {
		fmt.Println("\t getCachedAccessToken - cached access token does not exist")
		return ""
	}

	jsonFile, err := os.Open(getCachedAccessTokenFileUri())
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
	fmt.Println("createNewAccessToken")
	jwt := auth.GetServiceAccountJwt()
	tokenData := http.Auth(jwt)

	// persist token to configuration directory for caching purposes
	tokenDataJson, _ := json.Marshal(tokenData)
	if err := ioutil.WriteFile(getCachedAccessTokenFileUri(), tokenDataJson, 0700); err != nil {
		log.Fatal(err)
	}

	return tokenData.AccessToken
}

func getCachedAccessTokenFileUri() string {
	fmt.Println("getCachedAccessTokenFileUri")
	accessTokenFileUri := path.Join(utils.GetConfigDir(), CachedTokenFilename)
	return accessTokenFileUri
}

func existsCachedAccessToken() (bool, error) {
	fmt.Println("existsCachedAccessToken")
	accessTokenFileUri := getCachedAccessTokenFileUri()
	fmt.Println("\t existsCachedAccessToken - check: " + accessTokenFileUri)

	stat, err := os.Stat(accessTokenFileUri)
	if err == nil {
		if stat != nil {
			fmt.Println("\t\t existsCachedAccessToken - cached access token exists")
			return true, nil
		}
		return false, nil
	} else if os.IsNotExist(err) {
		fmt.Println("\t\t existsCachedAccessToken - cached access token does not exist")
		return false, nil
	}

	// e.g. permission denied
	return false, err
}
