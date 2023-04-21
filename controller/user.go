package controller

import (
	"net/http"

	"github.com/aldinofrizal/bumpaban/helpers"
	"github.com/aldinofrizal/bumpaban/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
}

func UserControllerImpl() *UserController {
	return &UserController{}
}

func (h *UserController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "successfully registered",
		"user":    user.GetResponse(),
	})
}

func (h *UserController) Login(ctx *gin.Context) {
	var login models.LoginRequest
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	result := models.DB.Where("email = ?", login.Email).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email/password"})
		return
	}

	validPassword := helpers.ComparePassword(login.Password, user.Password)
	if !validPassword {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email/password"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login OK",
		"user":    user.GetResponse(),
		"token":   helpers.GenerateJWT(jwt.MapClaims{"ID": user.ID}),
	})
}

func (h *UserController) Me(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"user":    user.GetResponse(),
	})
}
