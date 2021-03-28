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
	"strings"
	"yokanban-cli/internal/config"
	"yokanban-cli/internal/consts"
)

// HTTP the basic struct.
type HTTP struct {
	Client *http.Client
}

// Auth runs an authentication call to retrieve an access token for further api communication.
func (h *HTTP) Auth(jwt string) (TokenData, error) {
	apiURL := getAPIURL(consts.RouteOauthToken)
	data := url.Values{
		"grant_type": {"urn:ietf:params:oauth:grant-type:jwt-bearer"},
		"assertion":  {jwt},
	}

	req, _ := http.NewRequest(http.MethodPost, apiURL, strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	body, err := runHTTPCall(h.Client, req)
	if err != nil {
		return TokenData{}, err
	}

	var res TokenResponse

	if err := json.NewDecoder(strings.NewReader(body)).Decode(&res); err != nil {
		return TokenData{}, err
	}

	return res.Data, nil
}

// Delete runs a HTTP DELETE request to an API urlPath. Authentication is done via Bearer token.
func (h *HTTP) Delete(urlPath string, token string) (string, error) {
	apiURL := getAPIURL(urlPath)

	req, _ := http.NewRequest(http.MethodDelete, apiURL, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	return runHTTPCall(h.Client, req)
}

// Get runs a HTTP GET request to an API urlPath. Authentication is done via Bearer token.
func (h *HTTP) Get(urlPath string, token string) (string, error) {
	apiURL := getAPIURL(urlPath)

	req, _ := http.NewRequest(http.MethodGet, apiURL, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	return runHTTPCall(h.Client, req)
}

// Patch runs a HTTP PATCH request to an API urlPath. Authentication is done via Bearer token.
func (h *HTTP) Patch(urlPath string, token string, jsonBody string) (string, error) {
	apiURL := getAPIURL(urlPath)

	var jsonStr = []byte(jsonBody)

	req, _ := http.NewRequest(http.MethodPatch, apiURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	return runHTTPCall(h.Client, req)
}

// Post runs a HTTP POST request to an API urlPath. Authentication is done via Bearer token.
func (h *HTTP) Post(urlPath string, token string, jsonBody string) (string, error) {
	apiURL := getAPIURL(urlPath)

	var jsonStr = []byte(jsonBody)

	req, _ := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	return runHTTPCall(h.Client, req)
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
		log.Panic(err)
	}

	return string(body), nil
}
