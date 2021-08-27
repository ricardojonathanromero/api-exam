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
	reply := &models.Profit{}
	ctx, _ := initCtx()

	matchStage := bson.D{{"$match", bson.D{{"merchant_id", commerceID}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", nil}, {"total_amount", bson.D{{"$sum", "$fee"}}}}}}

	cursor, err := repo.transactionCol().Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		return 0, err
	}

	defer closeCursor(cursor)

	var totalAmount float32
	for cursor.Next(ctx) {
		err = cursor.Decode(&reply)
		if err == nil {
			totalAmount = float32(reply.TotalAmount)
			break
		}
	}

	return totalAmount, err
}

func (repo *repository) TotalAmount() (float32, error) {
	reply := &models.Profit{}
	ctx, _ := initCtx()

	groupStage := bson.D{{"$group", bson.D{{"_id", nil}, {"total_amount", bson.D{{"$sum", "$fee"}}}}}}

	cursor, err := repo.transactionCol().Aggregate(ctx, mongo.Pipeline{groupStage})
	if err != nil {
		return 0, err
	}

	defer closeCursor(cursor)
	var totalAmount float32
	for cursor.Next(ctx) {
		err = cursor.Decode(&reply)
		if err == nil {
			totalAmount = float32(reply.TotalAmount)
			break
		}
	}
	return totalAmount, err
}

func (repo *repository) transactionCol() *mongo.Collection {
	return repo.sess.Database("management").Collection("transactions")
}
