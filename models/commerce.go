package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CommerceReq struct {
	MerchantName string `json:"merchant_name"`
	Commission   int32  `json:"commission" validate:"min=1,max=100"`
}

type UpdateCommerce struct {
	MerchantName string `json:"merchant_name,omitempty"`
	Commission   int32  `json:"commission,omitempty" validate:"min=1,max=100"`
}

type Commerce struct {
	MerchantID   primitive.ObjectID `json:"merchant_id" bson:"merchant_id,omitempty"`
	MerchantName string             `json:"merchant_name" bson:"merchant_name"`
	Commission   int32              `json:"commission" bson:"commission"`
	CreatedAt    *time.Time         `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt    *time.Time         `json:"updated_at" bson:"updated_at"`
}
