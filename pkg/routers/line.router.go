package routers

import (
	"test_go_project/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func RouteLine(route *gin.RouterGroup, handler *handlers.LineHandler) {
	route.POST("/", handler.Webhook)
}
