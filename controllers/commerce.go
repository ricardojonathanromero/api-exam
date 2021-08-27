package controllers

import (
	"github.com/banwire/api-exam/models"
	"github.com/bostaurus/jsend"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"time"
)

const source = "commerce.api"

func (c *controller) AddCommerce(ctx *gin.Context) {
	payload := &models.CommerceReq{}

	if err := ctx.ShouldBind(&payload); err != nil {
		log.Println("error binding payload: ", err.Error())
		ctx.JSON(http.StatusBadRequest, jsend.Fail(&models.ErrRes{
			Code:    "payload_malformed",
			Message: "The request is malformed or has invalid params, please check it and retry more later.",
			Details: err.Error(),
			Source:  source,
		}))
		return
	}

	if err := validator.Validate(payload); err != nil {
		log.Println("error validating payload: ", err.Error())
		ctx.JSON(http.StatusBadRequest, jsend.Fail(&models.ErrRes{
			Code:    "invalid_payload",
			Message: "The request has invalid values, please check your request.",
			Details: err.Error(),
			Source:  source,
		}))
		return
	}

	now := time.Now()
	repository := c.getRepo()
	id := primitive.NewObjectID()

	err := repository.InsertCommerce(&models.Commerce{
		ID:           id,
		MerchantName: payload.MerchantName,
		Commission:   payload.Commission,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	})
	if err != nil {
		log.Println("error saving commerce: ", err.Error())
		ctx.JSON(http.StatusConflict, jsend.Fail(&models.ErrRes{
			Code:    "commerce_save_error",
			Message: "An error occurred trying to save your commerce in our systems",
			Details: err.Error(),
			Source:  source,
		}))
		return
	}

	ctx.JSON(http.StatusCreated, jsend.Success(map[string]string{
		"message":     "Commerce saved successfully",
		"merchant_id": id.Hex(),
	}))
	return
}

func (c *controller) UpdateCommerce(ctx *gin.Context) {
	payload := &models.UpdateCommerce{}
	commerceParam := ctx.Param("commerceID")

	if !primitive.IsValidObjectID(commerceParam) {
		log.Println("invalid commerce id: ", commerceParam)
		ctx.JSON(http.StatusBadRequest, jsend.Fail(&models.ErrRes{
			Code:    "invalid_commerce_id",
			Message: "The Commerce ID is not valid, please check your request",
			Details: "The commerce ID " + commerceParam + " is not valid",
			Source:  source,
		}))
		return
	}

	if err := ctx.ShouldBind(&payload); err != nil {
		log.Println("error binding payload: ", err.Error())
		ctx.JSON(http.StatusBadRequest, jsend.Fail(&models.ErrRes{
			Code:    "payload_malformed",
			Message: "The request is malformed or has invalid params, please check it and retry more later.",
			Details: err.Error(),
			Source:  source,
		}))
		return
	}

	if err := validator.Validate(payload); err != nil {
		log.Println("error validating payload: ", err.Error())
		ctx.JSON(http.StatusBadRequest, jsend.Fail(&models.ErrRes{
			Code:    "invalid_payload",
			Message: "The request has invalid values, please check your request.",
			Details: err.Error(),
			Source:  source,
		}))
		return
	}

	if payload.Commission <= 0 && len(payload.MerchantName) <= 0 {
		log.Println("both are empty!!!")
		ctx.JSON(http.StatusBadRequest, jsend.Fail(&models.ErrRes{
			Code:    "empty_payload",
			Message: "The request has invalid values, please check your request.",
			Details: "The request can not be empty.You need to send a valid value for one of params.",
			Source:  source,
		}))
		return
	}

	commerceID, _ := primitive.ObjectIDFromHex(commerceParam)
	now := time.Now()
	repository := c.getRepo()
	id := primitive.NewObjectID()

	count, err := repository.UpdateCommerce(commerceID, &models.Commerce{
		MerchantName: payload.MerchantName,
		Commission:   payload.Commission,
		UpdatedAt:    &now,
	})
	if err != nil {
		log.Println("error saving commerce: ", err.Error())
		ctx.JSON(http.StatusConflict, jsend.Fail(&models.ErrRes{
			Code:    "commerce_update_error",
			Message: "An error occurred trying to update your commerce in our systems",
			Details: err.Error(),
			Source:  source,
		}))
		return
	}

	if count <= 0 {
		log.Println("document no updated")
		ctx.JSON(http.StatusAccepted, jsend.Fail(map[string]string{
			"message":     "Commerce no updated because currently has the same values.",
			"merchant_id": id.Hex(),
		}))
		return
	}

	ctx.JSON(http.StatusOK, jsend.Success(map[string]string{
		"message":     "Commerce updated successfully",
		"merchant_id": id.Hex(),
	}))
	return
}
