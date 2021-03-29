package api

// AvatarDTO represents the exchange format of a single yokanban avatar.
type AvatarDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	IsVirtual bool   `json:"isVirtual"`
}
