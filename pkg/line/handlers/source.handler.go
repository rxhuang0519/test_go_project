package line

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type ISourceHandler interface {
	checkSource(ctx *gin.Context, id string) bool
	SaveSource(ctx *gin.Context, source *linebot.EventSource, receiveAt time.Time)
}
