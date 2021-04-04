package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"yokanban-cli/internal/accesstoken"
	"yokanban-cli/internal/consts"
	yohttp "yokanban-cli/internal/http"

	guuid "github.com/google/uuid"

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
}

// CreateBoardDTO represents the exchange format to create a single yokanban board.
type CreateBoardDTO struct {
	Name string `json:"name,omitempty"`
}

// UpdateBoardDTO represents the exchange format to update a single yokanban board.
type UpdateBoardDTO struct {
	NewName string `json:"newName,omitempty"`
}

// UserResponseDTO represents the exchange format of a user API response.
type UserResponseDTO struct {
	Success bool `json:"success"`
	Data    struct {
		*UserDTO
		Boards       []BoardDTO `json:"boards"`
		IsSocialUser bool       `json:"isSocialUser"`
	} `json:"data"`
}

// CreateBoard runs an API call to create a yokanban board.
func CreateBoard(model CreateBoardDTO) {
	log.Debugf("CreateBoard()")
	payload, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}
	body := runHTTPRequest(consts.RouteBoard, string(payload), requestOptions{retries: 0, maxRetries: 2, method: post})
	fmt.Println(body)
}

// CreateColumn runs an API call to create a column on a yokanban board.
func CreateColumn(boardID string, name string) {
	log.Debugf("CreateColumn()")
	uuid := guuid.New()
	column := ColumnEventDTO{
		Type:      "ADD",
		ElementID: uuid.String(),
		NewValues: &ColumnDTO{
			Type:  "COLUMN",
			Title: name,
			Shape: &ShapeDTO{
				X:      396.5,
				Y:      112.5,
				Width:  350,
				Height: 800,
			},
			Color:    "white",
			WipLimit: 0,
			IsLocked: false,
			ID:       uuid.String(),
			ZIndex:   1000,
		},
		SoftwareVersion: YoAPPVersion,
	}

	model := EventsContainerDTO{
		Event: EventDTO{
			Type:            "BULK",
			Events:          &[]ColumnEventDTO{column},
			SoftwareVersion: YoAPPVersion,
		},
	}

	payload, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}

	body := runHTTPRequest(path.Join(consts.RouteBoard, boardID), string(payload), requestOptions{retries: 0, maxRetries: 2, method: patch})
	fmt.Println(body)
}

// DeleteBoard runs an API call to delete a yokanban board.
func DeleteBoard(id string) {
	log.Debugf("DeleteBoard()")
	body := runHTTPRequest(path.Join(consts.RouteBoard, id), "", requestOptions{retries: 0, maxRetries: 2, method: delete})
	fmt.Println(body)
}

// GetBoard runs an API call to retrieve detail information of a yokanban board.
func GetBoard(id string) {
	log.Debugf("GetBoard()")
	body := runHTTPRequest(path.Join(consts.RouteBoard, id), "", requestOptions{retries: 0, maxRetries: 2, method: get})
	fmt.Println(body)
}

// UpdateBoard runs an API call to update a yokanban board.
func UpdateBoard(id string, model UpdateBoardDTO) {
	log.Debugf("UpdateBoard()")
	payload, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}
	// update the board name. Once more update possibilities have to be implemented, distinguish here.
	body := runHTTPRequest(path.Join(consts.RouteBoard, id, "name"), string(payload), requestOptions{retries: 0, maxRetries: 2, method: patch})
	fmt.Println(body)
}

// ListBoards runs an API call to retrieve a list of yokanban boards the current user has access to.
func ListBoards() {
	log.Debugf("ListBoards()")
	// for the list of boards the user has to be retrieved. Be aware that "user" scope is needed therefore!
	body := runHTTPRequest(consts.RouteUser, "", requestOptions{retries: 0, maxRetries: 2, method: get})

	// extract the boards
	var res UserResponseDTO
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		log.Fatal(err)
	}

	// create a boardList out of the response
	var boardList BoardList
	for _, b := range res.Data.Boards {
		board := &BoardShort{}
		board.Map(&b)
		boardList = append(boardList, board)
	}

	// generate the pretty printed output
	boardsPretty, err := json.MarshalIndent(boardList, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(boardsPretty))
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
