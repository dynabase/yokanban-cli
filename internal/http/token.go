package http

// TokenResponse represents a HTTP response of an Authentication call.
type TokenResponse struct {
	Data TokenData `json:"data"`
}

// TokenData represents the data of an Authentication call. It contains the access token including some metadata.
type TokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}
