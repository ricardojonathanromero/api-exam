package controllers

import (
	"github.com/banwire/api-exam/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IController interface {
	AddCommerce(ctx *gin.Context)
}

type controller struct {
	sess *mongo.Client
}

func (c *controller) getRepo() db.IRepository {
	return db.NewRepo(c.sess)
}

func (c *controller) calculateFee(commerceID primitive.ObjectID, amount float32) (int32, float32, error) {
	repository := c.getRepo()
	commerce, err := repository.GetCommerce(commerceID)
	if err != nil {
		return 0, 0, err
	}

	commission := float32(commerce.Commission) / 100
	fee := amount * commission

	return commerce.Commission, fee, nil
}

func NewController(sess *mongo.Client) *controller {
	return &controller{sess: sess}
}
