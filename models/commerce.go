package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type CommerceReq struct {
	Commission   int32  `json:"commission" validate:"min=1,max=100"`
	MerchantName string `json:"merchant_name" validate:"min=1,max=50,regexp=^([a-zA-Z&ñ\\s])+$"`
}

type UpdateCommerce struct {
	MerchantName string `json:"merchant_name,omitempty" validate:"min=1,max=50,regexp=^([a-zA-Z&ñ\\s])+$"`
	Commission   int32  `json:"commission,omitempty" validate:"min=1,max=100"`
}

type Commerce struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	MerchantName string             `json:"merchant_name" bson:"merchant_name"`
	Commission   int32              `json:"commission" bson:"commission"`
	CreatedAt    *time.Time         `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt    *time.Time         `json:"updated_at" bson:"updated_at"`
}
