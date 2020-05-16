package service

import (
	"fmt"
	"github.com/caglartelef/go-hello-world/model"
	"github.com/google/uuid"
)

// Token Service interface
type TokenService interface {
	// Create Token method received Create Token Request and returned Create Token Response
	CreateToken(request *model.CreateTokenRequest) *model.CreateTokenResponse
}

type TokenServiceImpl struct {
}

func NewTokenService() TokenService {
	return &TokenServiceImpl{}
}

func (t *TokenServiceImpl) CreateToken(request *model.CreateTokenRequest) *model.CreateTokenResponse {
	fmt.Println("Create Token Request, username:" + request.Username + " Password: " + request.Password)
	// We created fail case response
	response := &model.CreateTokenResponse{
		Success: false,
		Token:   "Can not created token! Username password is incorrect!",
	}

	if "caglar" == request.Username && "12345" == request.Password {
		response.Success = true
		response.Token = uuid.New().String()
		// We validated username and password and we returned success case response.
		return response
	}
	// We did not validate username and password and we returned fail case response.
	return response
}
