package api

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"yokanban-cli/internal/consts"
	"yokanban-cli/internal/http"
)

type requestParam struct {
	route      string
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
//func Create() {
//	// TODO to implement
//}
//
//func Delete() {
//	// TODO to implement
//}

// Test runs an API call to test current credentials
func Test() {
	log.Debug("Test()")
	body := runGetRequest(requestParam{route: consts.RouteOauthTest, retries: 0, maxRetries: 2})
	fmt.Println(body)
}

func runGetRequest(param requestParam) string {
	token := GetAccessToken()
	body, err := http.Get(param.route, token)
	if err != nil {
		if param.retries > param.maxRetries {
			log.Fatalf("Max retries of route %s reached", param.route)
		}

		// maybe token not valid anymore, create new one (will be cached for further requests)
		createNewAccessToken()

		// retry
		retries := param.retries + 1
		return runGetRequest(requestParam{route: param.route, retries: retries})
	}
	return body
}
