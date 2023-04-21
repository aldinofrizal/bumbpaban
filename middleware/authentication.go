package middleware

import (
	"net/http"

	"github.com/aldinofrizal/bumpaban/helpers"
	"github.com/aldinofrizal/bumpaban/models"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Please provide valid access token in Header Request",
			})
			return
		}

		claims, err := helpers.DecodeJWT(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Please provide valid access token in Header Request",
			})
			return
		}

		var loggedUser models.User
		result := models.DB.First(&loggedUser, claims["ID"])
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Please provide valid access token in Header Request",
			})
			return
		}

		ctx.Set("user", &loggedUser)
		ctx.Next()
	}
}
