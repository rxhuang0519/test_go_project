package handlers

import (
	"test_go_project/pkg/logger"
	"test_go_project/pkg/models"
	"test_go_project/pkg/services"
	"test_go_project/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageHandler struct {
	service *services.MessageService
}

func NewMessageHandler(db *mongo.Database) *MessageHandler {
	return &MessageHandler{
		service: services.NewMessageService(db),
	}
}
func (handler *MessageHandler) GetByMessageId(ctx *gin.Context) {
	msgId := ctx.Param("messageId")
	message, _ := handler.service.FindByMessageId(ctx, msgId)
	ctx.JSON(200, gin.H{"message": message})
	// msgId := ctx.Query("msgId")
}
func (handler *MessageHandler) GetByUserId(ctx *gin.Context) {
	usrId := ctx.Param("userId")
	messages, _ := handler.service.FindByUserId(ctx, usrId)
	ctx.JSON(200, gin.H{"messages": messages})
}
func (handler *MessageHandler) Get(ctx *gin.Context) {
	msgId := ctx.Query("id")
	usrId := ctx.Query("usr")
	input := &models.Message{
		MessageId: msgId,
		UserId:    usrId,
	}
	filter := handler.generateFilter(input)
	messages, _ := handler.service.Find(ctx, filter)
	ctx.JSON(200, gin.H{"messages": messages})
}
func (handler *MessageHandler) generateFilter(input *models.Message) interface{} {
	filter := utils.ToObject(input)
	logger.Debug.Printf("filter:%+v", filter)
	return filter
}
