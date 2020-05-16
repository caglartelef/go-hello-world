package controller

import (
	"github.com/caglartelef/go-hello-world/model"
	"github.com/caglartelef/go-hello-world/service"
	"github.com/gin-gonic/gin"
)

// We sould keep service in this struct
type TokenController struct {
	service service.TokenService
}

func NewTokenController(service service.TokenService) Controller {
	return &TokenController{service: service}
}

func (t *TokenController) Register(engine *gin.Engine) {
	// We should register method with request type in Gin Engine
	engine.GET("/createToken", t.CreateToken)
}

func (t *TokenController) CreateToken(ctx *gin.Context){
	var request model.CreateTokenRequest
	// We check to see if there is createTokenRequest in Gin-Gonics.
	err := ctx.Bind(&request)

	// if we did not find request type Gin-Gonics returns error.
	if err != nil {
		ctx.JSON(400, model.CreateTokenResponse{
			Success:	false,
			Token:		"Token not created! Request is invalid!",
		})
		return
	}
	// We call create token method from service.
	response:=t.service.CreateToken(&request)

	// If service' response failed, we returned fail response with HttpStatus code is 400
	if !response.Success {
		ctx.JSON(400, response)
		return
	}
	// We returned success response with HttpStatus code is 200
	ctx.JSON(200, response)
}