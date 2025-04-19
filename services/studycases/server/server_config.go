package server

import (
	"github.com/yujisoyama/go_microservices/pkg/utils"
)

type StudyCasesConfigs struct {
	port            string
	authManagerHost string
}

func (sc *StudyCases) SetConfigs() {
	sc.configs.port = utils.GetEnv("PORT")
	sc.configs.authManagerHost = utils.GetEnv("AUTHMANAGER_HOST")
}
