package routes

import (
	"github.com/aldinofrizal/bumpaban/controller"
	"github.com/aldinofrizal/bumpaban/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	user := r.Group("/users")
	userController := controller.UserControllerImpl()
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
		user.GET("/me", middleware.Authentication(), userController.Me)
	}
}
