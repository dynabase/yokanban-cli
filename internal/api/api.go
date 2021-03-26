package api

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"yokanban-cli/internal/accesstoken"
	"yokanban-cli/internal/consts"
	yohttp "yokanban-cli/internal/http"
)

// the HTTP request methods supported by yokanban
type requestMethod string

const (
	post  requestMethod = http.MethodPost
	get   requestMethod = http.MethodGet
	patch requestMethod = http.MethodPatch
)

type requestOptions struct {
	method     requestMethod
	retries    int
	maxRetries int
}

// CreateBoardModel describes all attributes of a board to be created.
type CreateBoardModel struct {
	Name string `json:"name,omitempty"`
}

// CreateBoard runs an API call to create a yokanban board.
func CreateBoard(model CreateBoardModel) {
	log.Debugf("CreateBoard()")
	payload, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}
	body := runHTTPRequest(consts.RouteBoard, string(payload), requestOptions{retries: 0, maxRetries: 2, method: post})
	fmt.Println(body)
}

// Test runs an API call to test current credentials
func Test() {
	log.Debug("Test()")
	body := runHTTPRequest(consts.RouteOauthTest, "", requestOptions{retries: 0, maxRetries: 2, method: get})
	fmt.Println(body)
}

func runHTTPRequest(route string, jsonBody string, options requestOptions) string {
	var body string
	var err error
	token := accesstoken.Get()

	switch method := options.method; method {
	case get:
		body, err = yohttp.Get(route, token)
	case post:
		body, err = yohttp.Post(route, token, jsonBody)
	case patch:
		body, err = yohttp.Patch(route, token, jsonBody)
	default:
		log.Fatalf("Method %s not implemented", method)
	}

	if err != nil {
		if options.retries > options.maxRetries {
			log.Fatalf("Max retries of route %s reached", route)
		}

		// maybe token not valid anymore, create new one (will be cached for further requests)
		accesstoken.Refresh()

		// retry
		retries := options.retries + 1
		return runHTTPRequest(route, jsonBody, requestOptions{retries: retries, maxRetries: options.maxRetries, method: options.method})
	}
	return body
}
