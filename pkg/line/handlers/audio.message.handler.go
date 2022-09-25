package line

import (
	"test_go_project/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

type AudioMessageHandler struct {
	*MessageHandler
}

func NewAudioMessageHandler(client *linebot.Client, db *mongo.Database) *AudioMessageHandler {
	return &AudioMessageHandler{
		MessageHandler: newMessageHandler(client, db),
	}
}
func (handler *AudioMessageHandler) SaveMessage(ctx *gin.Context, message linebot.Message, source *linebot.EventSource, receiveAt time.Time) {
	logger.Info.Println("Save Audio Message...")
	msg := message.(*linebot.AudioMessage)
	input := handler.newMessage(msg.ID, linebot.MessageTypeAudio, source, receiveAt)
	input.Audio = msg.ID
	handler.MessageHandler.saveMessage(ctx, input)
	logger.Info.Println("Save Audio Message Complete.")
}
func (handler *AudioMessageHandler) Reply(ctx *gin.Context, replyToken string) *linebot.ReplyMessageCall {
	logger.Info.Println("Reply Audio Message...: ", ctx.Keys["requestId"])
	return handler.client.ReplyMessage(replyToken, linebot.NewTextMessage("Recieve Audio."))
}
