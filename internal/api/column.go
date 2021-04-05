package api

import (
	"encoding/json"
	"path"
	"sort"
	"yokanban-cli/internal/consts"

	guuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const (
	defaultX      float32 = 396.5
	defaultY      float32 = 112.5
	defaultWidth  int     = 350
	defaultHeight int     = 800
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
func CreateColumn(boardID string, name string) string {
	log.Debugf("CreateColumn()")
	uuid := guuid.New()
	shape := getShapeDTO(boardID)
	column := ColumnEventDTO{
		Type:      "ADD",
		ElementID: uuid.String(),
		NewValues: &ColumnDTO{
			Type:     "COLUMN",
			Title:    name,
			Shape:    &shape,
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
	return body
}

func getShapeDTO(boardID string) ShapeDTO {
	// retrieve columns of board to calculate next X and Y position based on far right column
	boardDetails := GetBoard(boardID)
	cols := getColumns(boardDetails)

	shape := ShapeDTO{
		X:      defaultX,
		Y:      defaultY,
		Width:  defaultWidth,
		Height: defaultHeight,
	}

	if len(cols) == 0 {
		return shape
	}

	// sort by x value
	sort.Slice(cols, func(i, j int) bool {
		return cols[i].Shape.X < cols[j].Shape.X
	})

	lastCol := cols[len(cols)-1]

	shape.X = lastCol.Shape.X + float32(lastCol.Shape.Width)
	shape.Y = lastCol.Shape.Y
	shape.Width = lastCol.Shape.Width
	shape.Height = lastCol.Shape.Height

	return shape
}

func getColumns(boardDetails BoardDetails) []BoardElement {
	var columns []BoardElement
	for _, e := range boardDetails.Elements {
		if e.Type == "COLUMN" {
			columns = append(columns, e)
		}
	}
	return columns
}
