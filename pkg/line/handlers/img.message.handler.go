package line

import (
	"test_go_project/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

type ImgMessageHandler struct {
	*MessageHandler
}

func NewImgMessageHandler(client *linebot.Client, db *mongo.Database) *ImgMessageHandler {
	return &ImgMessageHandler{
		MessageHandler: newMessageHandler(client, db),
	}
}
func (handler *ImgMessageHandler) SaveMessage(ctx *gin.Context, message linebot.Message, source *linebot.EventSource, receiveAt time.Time) {
	logger.Info.Println("Save Img Message...")
	msg := message.(*linebot.ImageMessage)
	input := handler.newMessage(msg.ID, linebot.MessageTypeImage, source, receiveAt)
	input.Image = msg.ID
	handler.MessageHandler.saveMessage(ctx, input)
	logger.Info.Println("Save Img Message Complete.")
}
func (handler *ImgMessageHandler) Reply(ctx *gin.Context, replyToken string) *linebot.ReplyMessageCall {
	logger.Info.Println("Reply Image Message...: ", ctx.Keys["requestId"])
	return handler.client.ReplyMessage(replyToken, linebot.NewTextMessage("Recieve Image."))
}
