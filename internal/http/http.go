package http

import (
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

// Get runs a HTTP GET call to an API urlPath. Authentication is done via Bearer token.
func Get(urlPath string, token string) (string, error) {
	apiURL := getAPIURL(urlPath)

	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("Authorization", "Bearer "+token)
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

func getAPIURL(urlPath string) string {
	u, err := url.Parse(config.GetAPIURL())
	if err != nil {
		log.Fatal(err)
	}

	u.Path = path.Join(u.Path, urlPath)

	return u.String()
}
