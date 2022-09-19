/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"context"
	"test_go_project/cmd"
	"test_go_project/configs"
	"test_go_project/pkg/db"
	"test_go_project/pkg/line"
)

func main() {
	cmd.Execute()
	config := configs.Load()
	mongo_client, _ := db.Setup(context.Background(), &config)
	defer db.Disconnect(mongo_client)
	line.Setup(&config)
}
