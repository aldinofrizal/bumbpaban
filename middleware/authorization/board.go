package board

import (
	"net/http"

	"github.com/aldinofrizal/bumpaban/models"
	"github.com/gin-gonic/gin"
)

func ManageBoard() gin.HandlerFunc {
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

func CanAccess() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := ctx.MustGet("user").(*models.User)
		board := models.Board{}

		if err := models.DB.Preload("Members.User").First(&board, id).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"erorr": err.Error()})
			return
		}

		if !board.HasMember(int(user.ID)) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"erorr": "access forbidden"})
			return
		}

		ctx.Set("board", &board)
		ctx.Next()
	}
}
