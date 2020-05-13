package model

type CreateTokenResponse struct {
	Success bool   `json:success`
	Token   string `json:token`
}
