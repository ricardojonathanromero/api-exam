package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TransactionReq struct {
	Amount float32 `json:"amount" validate:"min=1,max=999999"`
}

type Transaction struct {
	TransactionID primitive.ObjectID `json:"transaction_id" bson:"transaction_id"`
	MerchantID    primitive.ObjectID `json:"merchant_id" bson:"merchant_id"`
	Amount        float32            `json:"amount" bson:"amount"`
	Commission    int32              `json:"commission" bson:"commission"`
	Fee           float32            `json:"fee" bson:"fee"`
	CreatedAt     *time.Time         `json:"created_at" bson:"created_at"`
}

type Profit struct {
	TotalAmount float32 `json:"total_amount" bson:"total_amount"`
}
