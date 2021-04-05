package api

import (
	"encoding/json"
	"fmt"
	"path"
	"yokanban-cli/internal/consts"

	log "github.com/sirupsen/logrus"
)

// BoardDTO represents the exchange format of a single yokanban board.
type BoardDTO struct {
	ID        string      `json:"_id"`
	Users     []UserDTO   `json:"users"`
	Name      string      `json:"name"`
	CreatedBy UserDTO     `json:"createdBy"`
	CreatedAt string      `json:"createdAt"`
	Avatars   []AvatarDTO `json:"avatars"`
}

// BoardShort represents the short version of a single yokanban board containing only the minimal dataset for an overview.
type BoardShort struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

// BoardList represents a list of yokanban boards.
type BoardList []*BoardShort

// Map maps a BoardDTO to a BoardShort.
func (b *BoardShort) Map(board *BoardDTO) {
	b.ID = board.ID
	b.Name = board.Name
	b.CreatedAt = board.CreatedAt
}

// CreateBoardDTO represents the exchange format to create a single yokanban board.
type CreateBoardDTO struct {
	Name string `json:"name,omitempty"`
}

// UpdateBoardDTO represents the exchange format to update a single yokanban board.
type UpdateBoardDTO struct {
	NewName string `json:"newName,omitempty"`
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
	boardList := BoardList{}
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
