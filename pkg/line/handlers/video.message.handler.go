package line

import (
	"test_go_project/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

type VideoMessageHandler struct {
	*MessageHandler
}

func NewVideoMessageHandler(client *linebot.Client, db *mongo.Database) *VideoMessageHandler {
	return &VideoMessageHandler{
		MessageHandler: newMessageHandler(client, db),
	}
}
func (handler *VideoMessageHandler) SaveMessage(ctx *gin.Context, message linebot.Message, source *linebot.EventSource, receiveAt time.Time) {
	logger.Info.Println("Save Video Message...")
	msg := message.(*linebot.VideoMessage)
	input := handler.newMessage(msg.ID, linebot.MessageTypeVideo, source, receiveAt)
	input.Video = msg.OriginalContentURL
	handler.MessageHandler.saveMessage(ctx, input)
	logger.Info.Println("Save Video Message Complete.")
}
func (handler *VideoMessageHandler) Reply(ctx *gin.Context, replyToken string) *linebot.ReplyMessageCall {
	logger.Info.Println("Reply Video Message...: ", ctx.Keys["requestId"])
	return handler.client.ReplyMessage(replyToken, linebot.NewTextMessage("Recieve Video."))
}
