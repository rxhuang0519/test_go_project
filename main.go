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
)

func main() {
	cmd.Execute()
	ctx := context.Background()
	cfg := configs.Load()
	// _cl, _ := line.Setup(cfg)
	client, _ := repository.Setup(ctx, cfg)
	defer repository.Disconnect(client)
	db := client.Database("demo")

	router := routers.NewRouter()
	lineClient := line.NewClient(cfg)
	lineHandler := handlers.NewLineHandler(cfg, lineClient, db)
	routers.RouteLine(router.Group("/"), lineHandler)
	router.Run()
}
