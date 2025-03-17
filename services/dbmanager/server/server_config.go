package server

import (
	"fmt"

	"github.com/yujisoyama/go_microservices/pkg/utils"
)

type DbManagerConfigs struct {
	port string
	apikey string
	dbHost string
	dbPort string
	dbUser string
	dbPass string
}

func (dbM *DbManager) SetConfigs() {
	dbM.configs.port = utils.GetEnv("PORT")
	dbM.configs.apikey = utils.GetEnv("API_KEY")
	dbM.configs.dbHost = utils.GetEnv("DB_HOST")
	dbM.configs.dbPort = utils.GetEnv("DB_PORT")
	dbM.configs.dbUser = utils.GetEnv("DB_USER")
	dbM.configs.dbPass = utils.GetEnv("DB_PASS")
}

func (dbM *DbManager) DbConnectString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/", dbM.configs.dbUser, dbM.configs.dbPass, dbM.configs.dbHost, dbM.configs.dbPort)
}
