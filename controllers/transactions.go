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

func (c *controller) AddTransaction(ctx *gin.Context) {
	payload := &models.TransactionReq{}

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

	now := time.Now()
	repository := c.getRepo()
	commerceID, _ := primitive.ObjectIDFromHex(commerceParam)
	commission, fee, err := c.calculateFee(commerceID, payload.Amount)
	if err != nil {
		log.Println("error saving transaction: ", err.Error())
		ctx.JSON(http.StatusConflict, jsend.Fail(&models.ErrRes{
			Code:    "fee_calculate_error",
			Message: "An error occurred trying to calculate fee",
			Details: err.Error(),
			Source:  source,
		}))
	}

	doc := &models.Transaction{
		TransactionID: primitive.NewObjectID(),
		MerchantID:    commerceID,
		Amount:        payload.Amount,
		Commission:    commission,
		Fee:           fee,
		CreatedAt:     &now,
	}

	err = repository.InsertTransaction(doc)
	if err != nil {
		log.Println("error saving transaction: ", err.Error())
		ctx.JSON(http.StatusConflict, jsend.Fail(&models.ErrRes{
			Code:    "transaction_save_error",
			Message: "An error occurred trying to save your commerce in our systems",
			Details: err.Error(),
			Source:  source,
		}))
		return
	}

	ctx.JSON(http.StatusCreated, jsend.Success(&doc))
	return
}

func (c *controller) GetProfit(ctx *gin.Context) {
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

	repository := c.getRepo()
	commerceID, _ := primitive.ObjectIDFromHex(commerceParam)

	total, err := repository.TotalAmountByCommerce(commerceID)
	if err != nil {
		log.Println("error calculating total amount: ", err.Error())
		ctx.JSON(http.StatusConflict, jsend.Fail(&models.ErrRes{
			Code:    "commerce_profits_error",
			Message: "An error occurred trying to calculate total profits",
			Details: err.Error(),
			Source:  source,
		}))
		return
	}

	ctx.JSON(http.StatusOK, jsend.Success(map[string]interface{}{
		"profits":     total,
		"merchant_id": commerceParam,
	}))
	return
}

func (c *controller) GetProfits(ctx *gin.Context) {
	repository := c.getRepo()

	total, err := repository.TotalAmount()
	if err != nil {
		log.Println("error calculating total amount: ", err.Error())
		ctx.JSON(http.StatusConflict, jsend.Fail(&models.ErrRes{
			Code:    "profits_error",
			Message: "An error occurred trying to calculate total profits",
			Details: err.Error(),
			Source:  source,
		}))
		return
	}

	ctx.JSON(http.StatusOK, jsend.Success(map[string]interface{}{
		"profits": total,
	}))
	return
}
