package api

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"yokanban-cli/internal/consts"
	"yokanban-cli/internal/elements"
	"yokanban-cli/internal/http"
)

type requestOptions struct {
	retries    int
	maxRetries int
}

//func List() {
//	// TODO to implement
//}
//
//func Get() {
//	// TODO to implement
//}
//
//func Delete() {
//	// TODO to implement
//}

// Create runs an API call to create a yokanban resource.
func Create(ele elements.YoElement) {
	log.Debugf("Create(%s)", ele)
	if ele == elements.Board {
		body := runPostRequest(consts.RouteBoard, "{}", requestOptions{retries: 0, maxRetries: 2})
		fmt.Println(body)
	} else {
		log.Fatalf("Creating of %s not implemented yet", ele)
	}
}

// Test runs an API call to test current credentials
func Test() {
	log.Debug("Test()")
	body := runGetRequest(consts.RouteOauthTest, requestOptions{retries: 0, maxRetries: 2})
	fmt.Println(body)
}

func runGetRequest(route string, options requestOptions) string {
	token := GetAccessToken()
	body, err := http.Get(route, token)
	if err != nil {
		if options.retries > options.maxRetries {
			log.Fatalf("Max retries of route %s reached", route)
		}

		// maybe token not valid anymore, create new one (will be cached for further requests)
		createNewAccessToken()

		// retry
		retries := options.retries + 1
		return runGetRequest(route, requestOptions{retries: retries, maxRetries: options.maxRetries})
	}
	return body
}

func runPostRequest(route string, jsonBody string, options requestOptions) string {
	token := GetAccessToken()
	body, err := http.Post(route, token, jsonBody)
	if err != nil {
		if options.retries > options.maxRetries {
			log.Fatalf("Max retries of route %s reached", route)
		}

		// maybe token not valid anymore, create new one (will be cached for further requests)
		createNewAccessToken()

		// retry
		retries := options.retries + 1
		return runPostRequest(route, jsonBody, requestOptions{retries: retries, maxRetries: options.maxRetries})
	}
	return body
}
