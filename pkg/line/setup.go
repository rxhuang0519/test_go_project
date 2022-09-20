package line

import (
	"test_go_project/configs"
	"test_go_project/pkg/logger"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Setup(config *configs.Config) (*linebot.Client, error) {
	logger.Info.Println("Setup LineBot...")
	bot, err := linebot.New(config.LINE_CHANNEL_SECRET, config.LINE_CHANNEL_TOKEN)
	if err != nil {
		logger.Error.Panicln("Setup linebot failed: \n", err)
	}
	logger.Info.Println("Setup LineBot complete.")
	return bot, err
}
