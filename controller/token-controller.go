package controller

import (
	"github.com/caglartelef/go-hello-world/model"
	"github.com/caglartelef/go-hello-world/service"
	"github.com/gin-gonic/gin"
)

type TokenController struct {
	service service.TokenService
}


func NewTokenController(service service.TokenService) Controller {
	return &TokenController{service: service}
}

func (t *TokenController) Register(engine *gin.Engine) {
	engine.GET("/createToken", t.CreateToken)
}

func (t *TokenController) CreateToken(ctx *gin.Context){
	var request model.CreateTokenRequest
	err := ctx.Bind(&request)

	if err != nil {
		ctx.JSON(400, model.CreateTokenResponse{
			Success:      false,
			Token:         "Token not created! Request is invalid!",
		})
		return
	}
	response:=t.service.CreateToken(&request)

	if !response.Success {
		ctx.JSON(400, response)
		return
	}
	ctx.JSON(200, response)
}