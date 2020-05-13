package main

import (
	"fmt"
	_ "github.com/caglartelef/go-hello-world/docs"
	"github.com/caglartelef/go-hello-world/startup"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// @title Go Project
// @version 1.0
// @description Go Project
// @termsOfService http://swagger.io/terms/

// @contact.name Caglar Telef
// @contact.url http://www.caglartelef.com
// @contact.email bilgi@caglartelef.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
//go:generate swag init
func main() {
	fmt.Println("Go Hello World!")
	router := gin.Default()
	router.RouterGroup.Handlers = router.RouterGroup.Handlers[0:0]
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	controllerArray := startup.Initialize()

	for _, key := range *controllerArray {
		key.Register(router)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/healthcheck", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})
	router.Run()
}
