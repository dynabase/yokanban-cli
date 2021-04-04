package api

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
