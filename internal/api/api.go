package api

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"yokanban-cli/internal/consts"
	"yokanban-cli/internal/http"
)

type requestMethod string

const (
	post  requestMethod = "post"
	get   requestMethod = "get"
	patch requestMethod = "patch"
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
	token := GetAccessToken()

	switch method := options.method; method {
	case get:
		body, err = http.Get(route, token)
	case post:
		body, err = http.Post(route, token, jsonBody)
	case patch:
		body, err = http.Patch(route, token, jsonBody)
	default:
		log.Fatalf("Method %s not implemented", method)
	}

	if err != nil {
		if options.retries > options.maxRetries {
			log.Fatalf("Max retries of route %s reached", route)
		}

		// maybe token not valid anymore, create new one (will be cached for further requests)
		createNewAccessToken()

		// retry
		retries := options.retries + 1
		return runHTTPRequest(route, jsonBody, requestOptions{retries: retries, maxRetries: options.maxRetries, method: options.method})
	}
	return body
}
