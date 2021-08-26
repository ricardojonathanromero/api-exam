package db

import (
	"github.com/banwire/api-exam/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *repository) InsertTransaction(transaction *models.Transaction) error {
	ctx, _ := initCtx()

	_, err := repo.transactionCol().InsertOne(ctx, &transaction)
	return err
}

func (repo *repository) TotalAmountByCommerce(commerceID primitive.ObjectID) (float32, error) {
	reply := make([]*models.Profit, 0)
	ctx, _ := initCtx()

	cursor, err := repo.transactionCol().Aggregate(ctx, mongo.Pipeline{
		bson.D{{"$match", bson.D{{"merchant_id", commerceID}}}},
		bson.D{{"$group", bson.D{{"_id", "$_id"}, {"total_amount", bson.D{{"$sum", "$fee"}}}}}},
	})
	if err != nil {
		return 0, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &reply)
	return reply[0].TotalAmount, err
}

func (repo *repository) TotalAmount() (float32, error) {
	reply := make([]*models.Profit, 0)
	ctx, _ := initCtx()

	cursor, err := repo.transactionCol().Aggregate(ctx, mongo.Pipeline{
		bson.D{{"$group", bson.D{{"_id", "$_id"}, {"total_amount", bson.D{{"$sum", "$fee"}}}}}},
	})
	if err != nil {
		return 0, err
	}

	defer cursor.Close(ctx)

	err = cursor.All(ctx, &reply)
	return reply[0].TotalAmount, err
}

func (repo *repository) transactionCol() *mongo.Collection {
	return repo.sess.Database("management").Collection("transactions")
}
