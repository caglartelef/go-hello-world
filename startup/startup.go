package startup

import (
	"github.com/caglartelef/go-hello-world/controller"
	"github.com/caglartelef/go-hello-world/service"
)

func Initialize() *[]controller.Controller {
	// Array for controllers with Controller Interface
	var list []controller.Controller

	// Token Service
	// This method returns TokenServiceImpl
	tokenService := service.NewTokenService()

	// We put Token Service in Token controller
	// This method returns TokenControllerImpl
	trackingController := controller.NewTokenController(tokenService)

	// We append list our controllers
	list = append(list, trackingController)
	return &list
}
