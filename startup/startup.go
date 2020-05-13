package startup

import (
	"github.com/caglartelef/go-hello-world/controller"
	"github.com/caglartelef/go-hello-world/service"
)

func Initialize() *[]controller.Controller {
	var list []controller.Controller
	tokenService := service.NewTokenService()
	trackingController := controller.NewTokenController(tokenService)
	list = append(list, trackingController)
	return &list
}
