package main

import (
	"context"

	"github.com/yujisoyama/go_microservices/services/studycases/server"
)

func main() {
	studyCases := server.NewStudyCases()
	err := studyCases.Run(context.Background())
	if err != nil {
		panic(err)
	}
}
