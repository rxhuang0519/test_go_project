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

type UserSourceHandler struct {
	client     *linebot.Client
	usrService *services.UserService
}

func NewUserSourceHandler(client *linebot.Client, db *mongo.Database) *UserSourceHandler {
	return &UserSourceHandler{
		client:     client,
		usrService: services.NewUserService(db),
	}
}
func (handler *UserSourceHandler) checkSource(ctx *gin.Context, id string) bool {
	logger.Info.Println("Current User Source: ", id)
	if usrs, _ := handler.usrService.FindByUserId(ctx, id); len(usrs) > 0 {
		return true
	} else {
		return false
	}
}

func (handler *UserSourceHandler) newSource(id string, receiveAt time.Time) *models.User {
	logger.Info.Println("new User Source...")
	usr := models.NewUser(id)
	usr.CreateAt = receiveAt
	usr.UpdateAt = receiveAt
	logger.Info.Println("new User Source Complete.")
	return usr
}
func (handler *UserSourceHandler) SaveSource(ctx *gin.Context, source *linebot.EventSource, receiveAt time.Time) {
	logger.Info.Println("Save User Source...: ", ctx.Keys["requestId"])
	if id := source.UserID; !handler.checkSource(ctx, id) {
		input := handler.newSource(id, receiveAt)
		handler.usrService.Create(ctx, input)
	}
	logger.Info.Println("Save User Source Complete.")
}
