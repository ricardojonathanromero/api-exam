package server

import (
	"fmt"
	"github.com/banwire/api-exam/routes"
	"github.com/banwire/api-exam/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
	"time"
)

func InitServerConfig(keys map[string]bool, sess *mongo.Client) *http.Server {
	port := utils.GetEnv("PORT", "3000")
	timeout, _ := time.ParseDuration(utils.GetEnv("TIMEOUT", "30"))
	writeTimeout, _ := time.ParseDuration(utils.GetEnv("W_TIMEOUT", "30"))
	maxHeadersBytes, _ := strconv.ParseInt(utils.GetEnv("MAX_BYTES", "100000"), 10, 64)
	addr := fmt.Sprintf(":%v", port)

	r := routes.NewRouter(keys, sess)

	return &http.Server{
		Addr:           addr,
		Handler:        r.Router(),
		ReadTimeout:    timeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: int(maxHeadersBytes),
	}
}
