package main

import (
	"context"

	"github.com/yujisoyama/go_microservices/services/dbmanager/server"
)

func main() {
	app := server.NewDbManager()
	app.Run(context.Background())
}