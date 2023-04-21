package board

import (
	"net/http"

	"github.com/aldinofrizal/bumpaban/models"
	"github.com/gin-gonic/gin"
)

func AddMember() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*models.User)
		board := models.Board{}

		if err := models.DB.Preload("Members").First(&board, ctx.Param("id")).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if int(user.ID) != board.GetOwnerId() {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden access"})
			return
		}

		ctx.Set("board", &board)
		ctx.Next()
	}
}
