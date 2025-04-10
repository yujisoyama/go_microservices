package server

import (
	"github.com/yujisoyama/go_microservices/pkg/utils"
)

type AuthManagerConfigs struct {
	port      string
	dbmHost   string
	dbmApikey string
}

func (am *AuthManager) SetConfigs() {
	am.configs.port = utils.GetEnv("PORT")
	am.configs.dbmHost = utils.GetEnv("DBM_HOST")
	am.configs.dbmApikey = utils.GetEnv("DBM_API_KEY")
}
