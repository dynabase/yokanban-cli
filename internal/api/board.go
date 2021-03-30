package api

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
