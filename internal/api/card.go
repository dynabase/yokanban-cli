package api

import (
	"encoding/json"
	"path"
	"yokanban-cli/internal/consts"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const (
	defaultCardX      float32 = 497.5
	defaultCardY      float32 = 438
	defaultCardWidth  int     = 250
	defaultCardHeight int     = 147
	defaultCardZIndex int     = 1000
)

// CardEventDTO represents the exchange format to create an event for a single yokanban card.
type CardEventDTO struct {
	Type            string   `json:"type"`
	ElementID       string   `json:"elementId"`
	OldValues       *CardDTO `json:"oldValues"`
	NewValues       *CardDTO `json:"newValues"`
	SoftwareVersion string   `json:"softwareVersion"`
}

// CardDTO represents the exchange format of a single yokanban card.
type CardDTO struct {
	Type       string    `json:"type"`
	Title      string    `json:"title"`
	Shape      *ShapeDTO `json:"shape"`
	Color      string    `json:"color"`
	IsArchived bool      `json:"isArchived"`
	ID         string    `json:"id"`
	ZIndex     int       `json:"zIndex"`
}

// CreateCardResponse represents the result of a create card API response.
type CreateCardResponse struct {
	Success bool                        `json:"success"`
	Data    []CreateCardResponseDetails `json:"data"`
}

// CreateCardResponseDetails represents the detail information of a create card API response.
type CreateCardResponseDetails struct {
	NewID string `json:"newId"`
	OldID string `json:"oldId"`
}

// CreateCard runs an API call to create a card on a yokanban board.
func (api *API) CreateCard(boardID string, name string, uuid uuid.UUID) []CreateCardResponseDetails {
	log.Debugf("CreateCard()")
	shape := ShapeDTO{
		X:      defaultCardX,
		Y:      defaultCardY,
		Width:  defaultCardWidth,
		Height: defaultCardHeight,
	}
	card := CardEventDTO{
		Type:      "ADD",
		ElementID: uuid.String(),
		NewValues: &CardDTO{
			Type:       "CARD",
			Title:      name,
			Shape:      &shape,
			Color:      "#FFFFFF",
			IsArchived: false,
			ID:         uuid.String(),
			ZIndex:     defaultCardZIndex,
		},
		SoftwareVersion: YoAPPVersion,
	}

	model := EventsContainerDTO{
		Event: EventDTO{
			Type:            "BULK",
			Events:          &[]CardEventDTO{card},
			SoftwareVersion: YoAPPVersion,
		},
	}

	payload, err := json.Marshal(model)
	if err != nil {
		log.Fatal(err)
	}

	body := api.runHTTPRequest(path.Join(consts.RouteBoard, boardID), string(payload), requestOptions{retries: 0, maxRetries: 2, method: patch})

	// extract the response data
	var res CreateCardResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		log.Fatal(err)
	}

	return res.Data
}
