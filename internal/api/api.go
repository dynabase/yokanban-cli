package api

import (
	"net/http"
	"yokanban-cli/internal/accesstoken"
	"yokanban-cli/internal/consts"
	yohttp "yokanban-cli/internal/http"

	log "github.com/sirupsen/logrus"
)

// the HTTP request methods supported by yokanban
type requestMethod string

// YoAPPVersion defines the semver version-string of the yokanban app. Set it for compatibility reasons.
const YoAPPVersion string = "0.4.12"

const (
	delete requestMethod = http.MethodDelete
	get    requestMethod = http.MethodGet
	patch  requestMethod = http.MethodPatch
	post   requestMethod = http.MethodPost
)

type requestOptions struct {
	method     requestMethod
	retries    int
	maxRetries int
	token      string
}

// API the basic struct.
type API struct {
	AccessToken accesstoken.TokenHandler
}

// Test runs an API call to test current credentials
func (api *API) Test() string {
	log.Debug("Test()")
	body := api.runHTTPRequest(consts.RouteOauthTest, "", requestOptions{retries: 0, maxRetries: 2, method: get})
	return body
}

func (api *API) runHTTPRequest(route string, jsonBody string, options requestOptions) string {
	var body string
	var err error

	token := options.token
	if token == "" {
		token = api.AccessToken.Get()
	}

	h := yohttp.HTTP{Client: &http.Client{}}

	switch method := options.method; method {
	case delete:
		body, err = h.Delete(route, token)
	case get:
		body, err = h.Get(route, token)
	case patch:
		body, err = h.Patch(route, token, jsonBody)
	case post:
		body, err = h.Post(route, token, jsonBody)
	default:
		log.Fatalf("Method %s not implemented", method)
	}

	if err != nil {
		if options.retries > options.maxRetries {
			log.Fatalf("Max retries of route %s reached", route)
		}

		// maybe token not valid anymore, create new one (will be cached for further requests)
		refreshToken := api.AccessToken.Refresh()

		// retry
		retries := options.retries + 1
		return api.runHTTPRequest(route, jsonBody, requestOptions{retries: retries, maxRetries: options.maxRetries, method: options.method, token: refreshToken})
	}
	return body
}
