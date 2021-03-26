package http

import (
	"bytes"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"yokanban-cli/internal/config"
	"yokanban-cli/internal/consts"
)

// Auth runs an authentication call to retrieve an access token for further api communication.
func Auth(jwt string) TokenData {
	apiURL := getAPIURL(consts.RouteOauthToken)
	data := url.Values{
		"grant_type": {"urn:ietf:params:oauth:grant-type:jwt-bearer"},
		"assertion":  {jwt},
	}

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		log.Fatal(err)
	}

	var res TokenResponse

	json.NewDecoder(resp.Body).Decode(&res)

	return res.Data
}

// Get runs a HTTP GET request to an API urlPath. Authentication is done via Bearer token.
func Get(urlPath string, token string) (string, error) {
	apiURL := getAPIURL(urlPath)

	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	return runHTTPCall(httpClient, req)
}

// Post runs a HTTP POST request to an API urlPath. Authentication is done via Bearer token.
func Post(urlPath string, token string, jsonBody string) (string, error) {
	apiURL := getAPIURL(urlPath)

	var jsonStr = []byte(jsonBody)

	httpClient := &http.Client{}
	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	return runHTTPCall(httpClient, req)
}

// Patch runs a HTTP PATCH request to an API urlPath. Authentication is done via Bearer token.
func Patch(urlPath string, token string, jsonBody string) (string, error) {
	apiURL := getAPIURL(urlPath)

	var jsonStr = []byte(jsonBody)

	httpClient := &http.Client{}
	req, _ := http.NewRequest("PATCH", apiURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	return runHTTPCall(httpClient, req)
}

func getAPIURL(urlPath string) string {
	u, err := url.Parse(config.GetAPIURL())
	if err != nil {
		log.Fatal(err)
	}

	u.Path = path.Join(u.Path, urlPath)

	return u.String()
}

func runHTTPCall(httpClient *http.Client, req *http.Request) (string, error) {
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return "", errors.New("API did not respond with expected status code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body), nil
}
