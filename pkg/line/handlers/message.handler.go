package line

import (
	"test_go_project/pkg/logger"
	"test_go_project/pkg/models"
	"test_go_project/pkg/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMessageHandler interface {
	SaveMessage(ctx *gin.Context, message linebot.Message, source *linebot.EventSource, receiveAt time.Time)
}
type MessageHandler struct {
	client     *linebot.Client
	msgService *services.MessageService
}

func newMessageHandler(client *linebot.Client, db *mongo.Database) *MessageHandler {
	return &MessageHandler{
		client:     client,
		msgService: services.NewMessageService(db),
	}
}
func (handler *MessageHandler) newMessage(id string, msgType linebot.MessageType, source *linebot.EventSource, receiveAt time.Time) *models.Message {
	logger.Info.Println("new Message...")
	msg := models.NewMessage(id)
	msg.CreateAt = receiveAt
	msg.UpdateAt = receiveAt
	msg.Type = string(msgType)
	msg.UserId = source.UserID
	msg.GroupId = source.GroupID
	msg.RoomId = source.RoomID
	logger.Info.Println("new Message Complete.")
	return msg
}
func (handler *MessageHandler) saveMessage(ctx *gin.Context, input *models.Message) {
	logger.Info.Println("Save Message...: ", ctx.Keys["requestId"])
	handler.msgService.Create(ctx, input)
	logger.Info.Println("Save Message Complete.")
}

// func (handler *MessageHandler) reply(ctx *gin.Context, replyToken string) *linebot.ReplyMessageCall {
// 	logger.Info.Println("Reply Message...: ", ctx.Keys["requestId"])
// 	return handler.client.ReplyMessage(replyToken, linebot.NewTextMessage("Recieved."))
// }
