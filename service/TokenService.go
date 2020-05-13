package service

import (
	"fmt"
	"github.com/caglartelef/go-hello-world/model"
	"github.com/google/uuid"
)

type TokenService interface {
	CreateToken(request *model.CreateTokenRequest) *model.CreateTokenResponse
}

type TokenServiceImpl struct {
}

func NewTokenService() TokenService {
	return &TokenServiceImpl{}
}

func (t *TokenServiceImpl) CreateToken(request *model.CreateTokenRequest) *model.CreateTokenResponse {
	fmt.Println("Create Token Request, username:" + request.Username + " Password: " + request.Password)
	response := &model.CreateTokenResponse{
		Success: false,
		Token:   "Can not created token! Username password is incorrect!",
	}

	if "caglar" == request.Username && "12345" == request.Password {
		response.Success = true
		response.Token = uuid.New().String()
		return response
	}
	return response
}
