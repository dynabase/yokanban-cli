package api

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"path"
	"yokanban-cli/internal/accesstoken"
	"yokanban-cli/internal/consts"
	yohttp "yokanban-cli/internal/http"
)

// the HTTP request methods supported by yokanban
type requestMethod string

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
}

// CreateBoardModel describes all attributes of a board to be created.
type CreateBoardModel struct {
	Name string `json:"name,omitempty"`
}

// DeleteBoardModel describes all attributes of a board to be deleted.
// a json representation is not needed since the id is not part of a HTTP request body.
type DeleteBoardModel struct {
	ID string
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

// DeleteBoard runs an API call to delete a yokanban board.
func DeleteBoard(model DeleteBoardModel) {
	log.Debugf("DeleteBoard()")
	body := runHTTPRequest(path.Join(consts.RouteBoard, model.ID), "", requestOptions{retries: 0, maxRetries: 2, method: delete})
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
		accesstoken.Refresh()

		// retry
		retries := options.retries + 1
		return runHTTPRequest(route, jsonBody, requestOptions{retries: retries, maxRetries: options.maxRetries, method: options.method})
	}
	return body
}
