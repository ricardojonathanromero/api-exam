package routes

import (
	"github.com/banwire/api-exam/controllers"
	"github.com/banwire/api-exam/middlewares"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type IRouter interface {
	Router() *gin.Engine
}

type router struct {
	keys map[string]bool
	sess *mongo.Client
}

func (r *router) Router() *gin.Engine {
	srv := gin.New()

	// define middlewares for trace
	srv.Use(gin.Logger())
	srv.Use(gin.Recovery())

	// define custom middleware for validate api key
	m := middlewares.AuthApiKey{ApiKeys: r.keys}
	auth := m.ApiKeyMiddle()

	// declare controller
	controller := controllers.NewController(r.sess)

	// configure protected API for expose to third-party
	ProtectedAPI := srv.Group("/api/v1").Use(auth)

	// commerce routes
	ProtectedAPI.POST("/commerce", controller.AddCommerce)
	ProtectedAPI.PUT("/commerce/:commerceID", controller.UpdateCommerce)

	// transactions routes
	ProtectedAPI.POST("/commerce/:commerceID/transactions", controller.AddTransaction)
	ProtectedAPI.GET("/commerce/:commerceID/profits", controller.GetProfit)
	ProtectedAPI.GET("/general/profits", controller.GetProfits)

	return srv
}

func NewRouter(keys map[string]bool, sess *mongo.Client) *router {
	return &router{keys: keys, sess: sess}
}
