package main

import (
	"context"

	"github.com/yujisoyama/go_microservices/services/authmanager/server"
)

func main() {
	app := server.NewAuthManager()
	err := app.Run(context.Background())
	if err != nil {
		panic(err)
	}
}