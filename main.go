/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"test_go_project/cmd"
	"test_go_project/configs"
	"test_go_project/pkg/line"
	"test_go_project/pkg/logger"
	"test_go_project/pkg/models"
	"test_go_project/pkg/repository"
	"test_go_project/pkg/services"
)

func main() {
	cmd.Execute()
	ctx := context.Background()
	cfg := configs.Load()
	line.Setup(cfg)
	client, _ := repository.Setup(ctx, cfg)
	defer repository.Disconnect(client)
	db := client.Database("demo")
	msgService := services.NewMessageService(db)
	msg := models.NewMessage("test message.")
	logger.Debug.Println("msgService:", msgService)
	logger.Debug.Println("msg:", msg)
	// msgService.FindById(ctx, "")
	// res, _ := msgService.UpdateById(ctx, "632f14b6f9b51cebcdfff903", &models.Message{Message: "updated message"})
	// res, _ := msgService.FindById(ctx, "632f14b6f9b51cebcdfff903")
	// logger.Info.Println(res.Id)
	res, _ := msgService.Create(ctx, msg)
	// res, _ := msgService.FindAll(ctx)
	// str, _ := json.Marshal(res)
	logger.Info.Printf("%+v", res)
	// _res, _ := msgService.FindAll(ctx)
	// logger.Info.Println(_res)
}
