package api

import (
	"encoding/json"
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

// GetBoardResponse represents the result of a get board API response.
type GetBoardResponse struct {
	Success bool         `json:"success"`
	Data    BoardDetails `json:"data"`
}

// BoardDetails represents the detail dataset of a yokanban board.
type BoardDetails struct {
	Avatars  []AvatarDTO    `json:"avatars"`
	Elements []BoardElement `json:"elements"`
}

// ListBoardsResponse represents the result of a list board API response.
type ListBoardsResponse struct {
	Success bool       `json:"success"`
	Data    []BoardDTO `json:"data"`
}

// BoardElement represents an element of a yokanban board.
type BoardElement struct {
	AssignedAvatars []string  `json:"assignedAvatars"`
	Blockers        []Blocker `json:"blockers"`
	Type            string    `json:"type"`
	Shape           ShapeDTO  `json:"shape"`
	Color           string    `json:"color"`
	Title           string    `json:"title"`
	IsArchived      bool      `json:"isArchived"`
	ZIndex          int       `json:"zIndex"`
	ID              string    `json:"id"`
}

// Blocker represents a blocker typically assigned to a card.
type Blocker struct {
	ID          string `json:"_id"`
	Description string `json:"description"`
	CreateAt    string `json:"createdAt"`
	CardID      string `json:"cardId"`
}

// CreateBoard runs an API call to create a yokanban board.
func (api *API) CreateBoard(model CreateBoardDTO) string {
	log.Debugf("CreateBoard()")
	payload, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}
	body := api.runHTTPRequest(consts.RouteBoard, string(payload), requestOptions{retries: 0, maxRetries: 2, method: post})
	return body
}

// DeleteBoard runs an API call to delete a yokanban board.
func (api *API) DeleteBoard(id string) string {
	log.Debugf("DeleteBoard()")
	body := api.runHTTPRequest(path.Join(consts.RouteBoard, id), "", requestOptions{retries: 0, maxRetries: 2, method: delete})
	return body
}

// GetBoard runs an API call to retrieve detail information of a yokanban board.
func (api *API) GetBoard(id string) BoardDetails {
	log.Debugf("GetBoard()")
	body := api.runHTTPRequest(path.Join(consts.RouteBoard, id), "", requestOptions{retries: 0, maxRetries: 2, method: get})

	// extract the board details
	var res GetBoardResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		log.Fatal(err)
	}

	return res.Data
}

// UpdateBoard runs an API call to update a yokanban board.
func (api *API) UpdateBoard(id string, model UpdateBoardDTO) string {
	log.Debugf("UpdateBoard()")
	payload, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}
	// update the board name. Once more update possibilities have to be implemented, distinguish here.
	body := api.runHTTPRequest(path.Join(consts.RouteBoard, id, "name"), string(payload), requestOptions{retries: 0, maxRetries: 2, method: patch})
	return body
}

// ListBoards runs an API call to retrieve a list of yokanban boards the current user has access to.
func (api *API) ListBoards() BoardList {
	log.Debugf("ListBoards()")
	// for the list of boards the user has to be retrieved. Be aware that "user" scope is needed therefore!
	body := api.runHTTPRequest(consts.RouteBoard, "", requestOptions{retries: 0, maxRetries: 2, method: get})

	// extract the boards
	var res ListBoardsResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		log.Fatal(err)
	}

	// create a boardList out of the response
	boardList := BoardList{}
	for _, b := range res.Data {
		board := &BoardShort{}
		board.Map(&b)
		boardList = append(boardList, board)
	}

	return boardList
}
