package routers

import (
	"test_go_project/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func RouteMessage(route *gin.RouterGroup, handler *handlers.MessageHandler) {
	route.GET("/:messageId", handler.GetByMessageId)
	route.GET("/usr/:userId", handler.GetByUserId)
	route.GET("/", handler.Get)
}
