package routes

import (
	"github.com/aldinofrizal/bumpaban/controller"
	"github.com/aldinofrizal/bumpaban/middleware"
	boardAuthz "github.com/aldinofrizal/bumpaban/middleware/authorization"
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

	private := r.Group("")
	private.Use(middleware.Authentication())

	board := private.Group("/boards")
	boardController := controller.BoardControllerImpl()
	{
		board.POST("", boardController.Create)
		board.GET("", boardController.Index)
		board.GET("/:id", boardAuthz.CanAccess(), boardController.Detail)
		board.DELETE("/:id", boardAuthz.ManageBoard(), boardController.Delete)
		board.POST("/users/:id", boardAuthz.ManageBoard(), boardController.AddMember)
		board.POST("/tasks/:id", boardAuthz.ManageBoard(), boardController.AddTask)
		board.PATCH("/tasks/:id/:task_id", boardAuthz.CanAccess(), boardController.UpdateStatus)
		board.DELETE("/tasks/:id/:task_id", boardAuthz.ManageBoard(), boardController.DeleteTask)
	}
}
