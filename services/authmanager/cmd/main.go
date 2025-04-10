package main

import (
	"context"

	"github.com/yujisoyama/go_microservices/services/authmanager/server"
)

func main() {
	authManager := server.NewAuthManager()
	err := authManager.Run(context.Background())
	if err != nil {
		panic(err)
	}
}