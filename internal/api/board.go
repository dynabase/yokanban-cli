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

// BoardShortDTO represents the exchange format of a single yokanban board containing only the data for an overview.
type BoardShortDTO struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

// BoardListDTO represents a list of yokanban boards.
type BoardListDTO []*BoardShortDTO

// Map maps a BoardDTO to a BoardShortDTO.
func (b *BoardShortDTO) Map(board *BoardDTO) {
	b.ID = board.ID
	b.Name = board.Name
	b.CreatedAt = board.CreatedAt
}
