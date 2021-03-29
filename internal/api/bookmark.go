package api

// BookmarkDTO represents the exchange format of a single yokanban bookmark.
type BookmarkDTO struct {
	ID        string  `json:"_id"`
	BoardID   string  `json:"boardId"`
	ZoomLevel float32 `json:"zoomLevel"`
	Name      string  `json:"name"`
	Offset    struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"offset"`
	CreatedAt string `json:"createdAt"`
}
