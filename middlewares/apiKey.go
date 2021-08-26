package middlewares

import (
	"errors"
	"github.com/banwire/api-exam/models"
	"github.com/bostaurus/jsend"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthApiKey struct {
	ApiKeys map[string]bool `json:"api_keys"`
}

func (auth *AuthApiKey) ApiKeyMiddle() gin.HandlerFunc {
	return func(context *gin.Context) {
		apiKey := context.Request.Header.Get("x-api-key")
		if len(apiKey) <= 0 {
			context.JSON(http.StatusUnauthorized, jsend.Fail(&models.ErrRes{
				Code:    "Unauthorized",
				Message: "Invalid api key, please check your request.",
				Details: "You must send a api key.",
				Source:  "auth.api",
			}))
			context.Abort()
			return
		}

		// validate api key
		err := auth.validateApiKey(apiKey)
		if err != nil {
			context.JSON(http.StatusForbidden, jsend.Fail(&models.ErrRes{
				Code:    "Forbidden",
				Message: err.Error(),
				Details: "Your API Key does not have permissions for consume this API.",
				Source:  "auth.api",
			}))
			context.Abort()
			return
		}

		context.Next()
	}
}

func (auth *AuthApiKey) validateApiKey(apiKey string) error {
	if len(auth.ApiKeys) <= 0 {
		return nil
	}

	if _, ok := auth.ApiKeys[apiKey]; !ok {
		return errors.New("the API Key is not valid")
	}

	return nil
}
