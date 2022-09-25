package line

import (
	"test_go_project/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

type TextMessageHandler struct {
	*MessageHandler
}

func NewTextMessageHandler(client *linebot.Client, db *mongo.Database) *TextMessageHandler {
	return &TextMessageHandler{
		MessageHandler: newMessageHandler(client, db),
	}
}
func (handler *TextMessageHandler) SaveMessage(ctx *gin.Context, message linebot.Message, source *linebot.EventSource, receiveAt time.Time) {
	logger.Info.Println("Save Text Message...")
	msg := message.(*linebot.TextMessage)
	input := handler.newMessage(msg.ID, linebot.MessageTypeText, source, receiveAt)
	input.Text = msg.Text
	handler.MessageHandler.saveMessage(ctx, input)
	logger.Info.Println("Save Text Message Complete.")
}
func (handler *TextMessageHandler) Reply(ctx *gin.Context, replyToken string) *linebot.ReplyMessageCall {
	logger.Info.Println("Reply Text Message...: ", ctx.Keys["requestId"])
	return handler.client.ReplyMessage(replyToken, linebot.NewTextMessage("Recieve Text."))
}
