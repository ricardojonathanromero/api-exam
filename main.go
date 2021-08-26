package main

import (
	"github.com/banwire/api-exam/config"
	"github.com/banwire/api-exam/server"
	"github.com/banwire/api-exam/utils"
	"log"
)

func main() {
	// get api key value from env variable
	apiKey := utils.GetEnv("API_KEY", "")
	keys := utils.StringToMap(apiKey)

	// init db conn
	dbHost := utils.GetEnv("DB_HOST", "")
	if len(dbHost) <= 0 {
		log.Fatalf("you need to provide a host db for configure the conenction, the system could not retrieve the value: %v", dbHost)
	}
	sess, err := config.CreateSession(dbHost)
	if err != nil {
		log.Fatalf("error configuring db session: %v", err)
	}

	// configure server
	ws := server.InitServerConfig(keys, sess)

	// init server
	if err := ws.ListenAndServe(); err != nil {
		log.Fatalf("error init server: %v", err)
	}
}
