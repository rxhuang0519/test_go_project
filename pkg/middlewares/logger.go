package middlewares

import (
	"test_go_project/pkg/logger"

	"github.com/gin-gonic/gin"
)

func RequestLogger(ctx *gin.Context) {
	logger.Info.Printf("%s[%s] %s%s: %s", logger.Blue(), ctx.Keys["requestId"], logger.Green(), ctx.Request.Method, ctx.FullPath())
}
func Response(ctx *gin.Context) {
	logger.Info.Println(logger.Green(), "Path: ", ctx.FullPath())
	logger.Info.Println(logger.Green(), "Header: ", ctx.Request.Header)
	logger.Info.Println(logger.Green(), "Method: ", ctx.Request.Method)
	logger.Info.Println(logger.Green(), "Body: ", ctx.Request.Body)
}
