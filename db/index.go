package db

import (
	"context"
	"github.com/banwire/api-exam/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type IRepository interface {
	GetCommerce(id primitive.ObjectID) (*models.Commerce, error)
	InsertCommerce(commerce *models.Commerce) error
	InsertTransaction(transaction *models.Transaction) error
	TotalAmount() (float32, error)
	TotalAmountByCommerce(commerceID primitive.ObjectID) (float32, error)
	UpdateCommerce(id primitive.ObjectID, commerce *models.Commerce) (int32, error)
}

type repository struct {
	sess *mongo.Client
}

func initCtx() (context.Context, context.CancelFunc) {
	ctx, c := context.WithTimeout(context.TODO(), 10*time.Second)
	return ctx, c
}

func NewRepo(sess *mongo.Client) *repository {
	return &repository{sess: sess}
}
