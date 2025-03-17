package main

import (
	"context"

	"github.com/yujisoyama/go_microservices/services/dbmanager/server"
)

func main() {
	app := server.NewDbManager()
	err := app.Run(context.Background())
	if err != nil {
		panic(err)
	}
}