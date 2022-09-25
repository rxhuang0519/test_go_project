/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"test_go_project/cmd"
	"test_go_project/configs"
	"test_go_project/pkg/handlers"
	"test_go_project/pkg/line"
	"test_go_project/pkg/repository"
	"test_go_project/pkg/routers"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	dbClient   *mongo.Client
	lineClient *linebot.Client
	rootRouter *gin.Engine
)

func main() {
	defer repository.Disconnect(dbClient)
	start("demo")
}
func init() {
	cmd.Execute()
	cfg := configs.Load()
	ctx := context.Background()
	dbClient = repository.NewClient(ctx, cfg)
	lineClient = line.NewClient(cfg)
	rootRouter = routers.NewRouter()
}
func start(dbName string) {
	db := dbClient.Database(dbName)
	lineHandler := handlers.NewLineHandler(lineClient, db)
	msgHandler := handlers.NewMessageHandler(db)
	routers.RouteLine(rootRouter.Group("/"), lineHandler)
	routers.RouteMessage(rootRouter.Group("/messages"), msgHandler)
	rootRouter.Run()
}
