package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"yokanban-cli/internal/config"
)

func Auth(jwt string) TokenData {
	apiUrl := getApiUrl("/auth/oauth2/token")
	data := url.Values{
		"grant_type": {"urn:ietf:params:oauth:grant-type:jwt-bearer"},
		"assertion":  {jwt},
	}

	resp, err := http.PostForm(apiUrl, data)
	if err != nil {
		log.Fatal(err)
	}

	var res TokenResponse

	json.NewDecoder(resp.Body).Decode(&res)

	return res.Data
}

func Get(urlPath string, token string) string {
	apiUrl := getApiUrl(urlPath)

	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Fatal("API did not respond with expected status code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func getApiUrl(urlPath string) string {
	u, err := url.Parse(config.GetApiUrl())
	if err != nil {
		log.Fatal(err)
	}

	u.Path = path.Join(u.Path, urlPath)

	return u.String()
}
