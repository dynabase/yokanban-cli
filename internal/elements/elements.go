package elements

// YoElement defines yokanban elements which can be processed by yokanban-cli.
type YoElement string

// the YoElement enums
const (
	Board  YoElement = "board"
	Boards YoElement = "boards"
	Column YoElement = "column"
	Card   YoElement = "card"
)
