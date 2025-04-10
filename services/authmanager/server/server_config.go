package server

import (
	"github.com/yujisoyama/go_microservices/pkg/utils"
)

type AuthManagerConfigs struct {
	port      string
	dbmApikey string
}

func (am *AuthManager) SetConfigs() {
	am.Configs.port = utils.GetEnv("PORT")
	am.Configs.dbmApikey = utils.GetEnv("DBM_API_KEY")
}
