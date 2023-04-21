package controller

import (
	"net/http"

	"github.com/aldinofrizal/bumpaban/models"
	"github.com/aldinofrizal/bumpaban/services"
	"github.com/gin-gonic/gin"
)

type BoardController struct{}

func BoardControllerImpl() *BoardController {
	return &BoardController{}
}

func (r *BoardController) Create(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	var board models.Board
	if err := ctx.ShouldBindJSON(&board); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateBoard(&board, user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"board":   board,
	})
}

func (r *BoardController) AddMember(ctx *gin.Context) {
	board := ctx.MustGet("board").(*models.Board)
	addUser := struct {
		UserId int `form:"user_id" json:"user_id" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&addUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.AddMember(board, addUser.UserId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success add member",
	})
}

func (r *BoardController) Index(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	if err := models.DB.Model(user).Preload("Boards.Members.User").Find(user).Error; err != nil {
		panic(err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"boards":  services.FormatBoards(user.Boards),
	})
}

func (r *BoardController) Detail(ctx *gin.Context) {
	id := ctx.Param("id")
	board := models.Board{}

	if err := models.DB.Preload("Members.User").First(&board, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erorr": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"board":   board.GetIndexResponse(),
	})
}

func (r *BoardController) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
