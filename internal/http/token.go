package http

type TokenResponse struct {
	Data TokenData `json:"data"`
}

type TokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}
