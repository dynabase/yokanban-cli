package api

import (
	"encoding/json"
	"fmt"
	"path"
	"yokanban-cli/internal/consts"

	guuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// ColumnEventDTO represents the exchange format to create an event for a single yokanban column.
type ColumnEventDTO struct {
	Type            string     `json:"type"`
	ElementID       string     `json:"elementId"`
	OldValues       *ColumnDTO `json:"oldValues"`
	NewValues       *ColumnDTO `json:"newValues"`
	SoftwareVersion string     `json:"softwareVersion"`
}

// ColumnDTO represents the exchange format of a single yokanban column.
type ColumnDTO struct {
	Type     string    `json:"type"`
	Title    string    `json:"title"`
	Shape    *ShapeDTO `json:"shape"`
	Color    string    `json:"color"`
	WipLimit uint      `json:"wipLimit"`
	IsLocked bool      `json:"isLocked"`
	ID       string    `json:"id"`
	ZIndex   int       `json:"zIndex"`
}

// ShapeDTO represents the exchange format of a column shape.
type ShapeDTO struct {
	X      float32 `json:"x"`
	Y      float32 `json:"y"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
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
