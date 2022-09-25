package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MetaData(ctx *gin.Context) {
	setRequestId(ctx)
	ctx.Next()
}
func setRequestId(ctx *gin.Context) {
	requestId := uuid.New().String()
	ctx.Set("requestId", requestId)
}
