package line

import (
	"test_go_project/configs"
	"test_go_project/pkg/logger"

	"github.com/line/line-bot-sdk-go/linebot"
)

func NewClient(config *configs.Config) *linebot.Client {
	logger.Info.Println("Setup LineBot...")
	client, err := linebot.New(config.LINE_CHANNEL_SECRET, config.LINE_CHANNEL_TOKEN)
	if err != nil {
		logger.Error.Panicln("Setup linebot failed: \n", err)
	}
	logger.Info.Println("Setup LineBot complete.")
	return client
}
