package db

import (
	"github.com/banwire/api-exam/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *repository) GetCommerce(id primitive.ObjectID) (*models.Commerce, error) {
	reply := &models.Commerce{}
	ctx, _ := initCtx()

	err := repo.commerceCol().FindOne(ctx, bson.M{"_id": id}).Decode(&reply)
	return reply, err
}

func (repo *repository) InsertCommerce(commerce *models.Commerce) error {
	ctx, _ := initCtx()

	_, err := repo.commerceCol().InsertOne(ctx, &commerce)
	return err
}

func (repo *repository) UpdateCommerce(id primitive.ObjectID, commerce *models.Commerce) (int32, error) {
	ctx, _ := initCtx()

	result, err := repo.commerceCol().UpdateByID(ctx, id, bson.M{"$set": &commerce})
	if err != nil {
		return 0, err
	}

	return int32(int(result.ModifiedCount)), err
}

func (repo *repository) commerceCol() *mongo.Collection {
	return repo.sess.Database("management").Collection("commerces")
}
