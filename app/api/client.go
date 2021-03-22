package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"yokanban-cli/app"
)

type TokenResponse struct {
	Data TokenData `json:"data"`
}

type TokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func Auth(jwt string) string {
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

	return res.Data.AccessToken
}

func Get(urlPath string, token string) {
	apiUrl := getApiUrl(urlPath)

	httpClient := &http.Client{}
	req, _ := http.NewRequest("GET", apiUrl, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func getApiUrl(urlPath string) string {
	u, err := url.Parse(app.GetApiUrl())
	if err != nil {
		log.Fatal(err)
	}

	u.Path = path.Join(u.Path, urlPath)

	return u.String()
}
