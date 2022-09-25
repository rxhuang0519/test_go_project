package line

import (
	"test_go_project/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileMessageHandler struct {
	*MessageHandler
}

func NewFileMessageHandler(client *linebot.Client, db *mongo.Database) *FileMessageHandler {
	return &FileMessageHandler{
		MessageHandler: newMessageHandler(client, db),
	}
}
func (handler *FileMessageHandler) SaveMessage(ctx *gin.Context, message linebot.Message, source *linebot.EventSource, receiveAt time.Time) {
	logger.Info.Println("Save File Message...")
	msg := message.(*linebot.FileMessage)
	input := handler.newMessage(msg.ID, linebot.MessageTypeFile, source, receiveAt)
	input.File = msg.FileName
	handler.MessageHandler.saveMessage(ctx, input)
	logger.Info.Println("Save File Message Complete.")
}
func (handler *FileMessageHandler) Reply(ctx *gin.Context, replyToken string) *linebot.ReplyMessageCall {
	logger.Info.Println("Reply File Message...: ", ctx.Keys["requestId"])
	return handler.client.ReplyMessage(replyToken, linebot.NewTextMessage("Recieve File."))
}
