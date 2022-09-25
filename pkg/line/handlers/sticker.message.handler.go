package line

import (
	"test_go_project/pkg/logger"
	"test_go_project/pkg/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

type StickerMessageHandler struct {
	*MessageHandler
}

func NewStickerMessageHandler(client *linebot.Client, db *mongo.Database) *StickerMessageHandler {
	return &StickerMessageHandler{
		MessageHandler: newMessageHandler(client, db),
	}
}
func (handler *StickerMessageHandler) newSticker(msg *linebot.StickerMessage) *models.Sticker {
	return &models.Sticker{
		StickerId:   msg.StickerID,
		StickerType: string(msg.StickerResourceType),
		PackageId:   msg.PackageID,
	}
}
func (handler *StickerMessageHandler) SaveMessage(ctx *gin.Context, message linebot.Message, source *linebot.EventSource, receiveAt time.Time) {
	logger.Info.Println("Save Sticker Message...")
	msg := message.(*linebot.StickerMessage)
	input := handler.newMessage(msg.ID, linebot.MessageTypeSticker, source, receiveAt)
	input.Sticker = handler.newSticker(msg)
	handler.MessageHandler.saveMessage(ctx, input)
	logger.Info.Println("Save Sticker Message Complete.")
}
