package model

// Create Token Response
type CreateTokenResponse struct {
	// Success type is bool
	Success bool `json:success`
	// Token type is string
	Token string `json:token`
}
