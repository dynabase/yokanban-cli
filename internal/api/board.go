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
