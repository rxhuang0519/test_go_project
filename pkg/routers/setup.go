package routers

import (
	"test_go_project/cmd"
	"test_go_project/pkg/logger"
	"test_go_project/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	logger.Info.Println("Setup Gin Router...")
	env := cmd.ENV()
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.Use(middlewares.MetaData, middlewares.RequestLogger)
	logger.Info.Println("Setup Gin Router complete.")
	return router
}
