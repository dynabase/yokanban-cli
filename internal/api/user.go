package api

// UserDTO represents the exchange format of a single yokanban user.
type UserDTO struct {
	ID          string        `json:"_id"`
	FirstName   string        `json:"firstName"`
	LastName    string        `json:"lastName"`
	Email       string        `json:"email"`
	AvatarURL   string        `json:"avatarUrl"`
	CreatedAt   string        `json:"createdAt"`
	Bookmarks   []BookmarkDTO `json:"bookmarks"`
	LastBoardID string        `json:"lastBoardId"`
}

// UserResponseDTO represents the exchange format of a user API response.
type UserResponseDTO struct {
	Success bool `json:"success"`
	Data    struct {
		*UserDTO
		Boards       []BoardDTO `json:"boards"`
		IsSocialUser bool       `json:"isSocialUser"`
	} `json:"data"`
}
