package handlers

import (
	"errors"
	line "test_go_project/pkg/line/handlers"
	"test_go_project/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

//	type ILineHandler interface {
//		Reply()
//	}
type LineHandler struct {
	// cfg    *configs.Config
	client *linebot.Client
	db     *mongo.Database
}

func NewLineHandler(client *linebot.Client, db *mongo.Database) *LineHandler {
	return &LineHandler{
		// cfg:    cfg,
		client: client,
		db:     db,
	}
}

//	func (handler *LineHandler) Auth(ctx *gin.Context) {
//		signature := ctx.Request.Header.Get("x-line-signature")
//		isVaild := false
//		/* Read Body */
//		if body, err := ctx.GetRawData(); err != nil {
//			logger.Error.Println("Read Error:", err)
//		} else {
//			/* Reset Reader */
//			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
//			/* Decoding */
//			if decoded, err := base64.StdEncoding.DecodeString(signature); err != nil {
//				logger.Error.Println("Decode Error:", err)
//			} else {
//				/* Expected hash */
//				hash := hmac.New(sha256.New, []byte(handler.cfg.LINE_CHANNEL_SECRET))
//				hash.Write(body)
//				isVaild = hmac.Equal(decoded, hash.Sum(nil))
//				logger.Info.Println("isVaild: ", isVaild)
//				logger.Info.Println("body:", string(body))
//			}
//		}
//		// return isVaild
//		if isVaild {
//			ctx.Next()
//		} else {
//			logger.Error.Println("Authentication failed:", signature)
//			ctx.AbortWithStatus(401)
//		}
//	}
func (handler *LineHandler) parseRequest(ctx *gin.Context) []*linebot.Event {
	logger.Info.Println("Line Parse")
	events, err := handler.client.ParseRequest(ctx.Request)
	if err != nil {
		logger.Error.Println("Error:", err)
		if err == linebot.ErrInvalidSignature {
			ctx.AbortWithStatus(401)
		}
		ctx.AbortWithError(500, errors.New("invalid request format"))
	}
	return events
}

func (handler *LineHandler) Webhook(ctx *gin.Context) {
	logger.Info.Println("Handle...")
	events := handler.parseRequest(ctx)
	for _, event := range events {
		handler.handleSource(ctx, event)
		switch event.Type {
		case linebot.EventTypeMessage:
			handler.handleMessage(ctx, event)
		default:
			logger.Info.Printf("Unsupported Event Type::%s:%+v", event.Type, event)
		}
	}
	ctx.Status(200)

}
func (handler *LineHandler) handleSource(ctx *gin.Context, event *linebot.Event) {
	logger.Info.Println("save Source...")
	var h line.ISourceHandler
	switch event.Source.Type {
	case linebot.EventSourceTypeUser:
		h = line.NewUserSourceHandler(handler.client, handler.db)
	default:
		logger.Info.Printf("Unknown Source::%s: %+v", event.Source.Type, event.Source)
	}
	h.SaveSource(ctx, event.Source, event.Timestamp)
	logger.Info.Println("save Source Complet.")

}
func (handler *LineHandler) handleMessage(ctx *gin.Context, event *linebot.Event) {
	logger.Info.Println("Handle Message...")
	var h line.IMessageHandler
	switch msg := event.Message.(type) {
	case *linebot.TextMessage:
		h = line.NewTextMessageHandler(handler.client, handler.db)
	case *linebot.ImageMessage:
		h = line.NewImgMessageHandler(handler.client, handler.db)
	case *linebot.StickerMessage:
		h = line.NewStickerMessageHandler(handler.client, handler.db)
	case *linebot.AudioMessage:
		h = line.NewAudioMessageHandler(handler.client, handler.db)
	case *linebot.VideoMessage:
		h = line.NewVideoMessageHandler(handler.client, handler.db)
	case *linebot.FileMessage:
		h = line.NewFileMessageHandler(handler.client, handler.db)
	default:
		logger.Info.Printf("Unknown Message::%T: %+v", msg, msg)
	}
	h.SaveMessage(ctx, event.Message, event.Source, event.Timestamp)
	if res, err := h.Reply(ctx, event.ReplyToken).Do(); err != nil {
		ctx.AbortWithError(500, err)
	} else {
		logger.Info.Println("Handle Message Response RequestID:", res.RequestID)
	}
	logger.Info.Println("Handle Message Complete.")

}
