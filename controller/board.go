package controller

import (
	"fmt"
	"net/http"

	"github.com/aldinofrizal/bumpaban/models"
	"github.com/aldinofrizal/bumpaban/services"
	"github.com/gin-gonic/gin"
)

type BoardController struct{}

type UpdateStatusRequest struct {
	Status int `form:"status" json:"status" binding:"required"`
}

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
		"board":   board.GetIndexResponse(),
	})
}

func (r *BoardController) AddMember(ctx *gin.Context) {
	board := ctx.MustGet("board").(*models.Board)
	addUser := struct {
		Email string `form:"email" json:"email" binding:"required"`
	}{}

	if err := ctx.ShouldBindJSON(&addUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("email = ?", addUser.Email).First(&user).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := services.AddMember(board, int(user.ID)); err != nil {
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
	board := ctx.MustGet("board").(*models.Board)
	user := ctx.MustGet("user").(*models.User)

	result := models.DB.Where("board_id = ?", board.ID).Preload("Assignee").Find(&board.Tasks)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":             "OK",
		"board":               board.GetIndexResponse(),
		"task_status_mapping": models.STATUS,
		"is_owner":            board.GetOwnerId() == int(user.ID),
	})
}

func (r *BoardController) Delete(ctx *gin.Context) {
	board := ctx.MustGet("board").(*models.Board)
	if err := models.DB.Delete(board).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Success delete %s board", board.Title),
	})
}

func (r *BoardController) AddTask(ctx *gin.Context) {
	board := ctx.MustGet("board").(*models.Board)
	requestTask := models.TaskRequest{
		BoardId: int(board.ID),
	}
	if err := ctx.ShouldBindJSON(&requestTask); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask := models.Task{
		Title:       requestTask.Title,
		Description: requestTask.Description,
		BoardId:     requestTask.BoardId,
		// AssigneeId:  requestTask.AssigneeId,
	}
	if err := models.DB.Create(&newTask).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":             "Successfully add new task",
		"task":                newTask,
		"task_status_mapping": models.STATUS,
	})
}

func (r *BoardController) UpdateStatus(ctx *gin.Context) {
	user := ctx.MustGet("user").(*models.User)
	var task models.Task
	if err := models.DB.First(&task, ctx.Param("task_id")).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var requestStatus UpdateStatusRequest
	if err := ctx.ShouldBindJSON(&requestStatus); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Status = requestStatus.Status
	if task.Unassigned() {
		task.AssigneeId = int(user.ID)
	}

	result := models.DB.Save(&task)
	if result.Error != nil {
		fmt.Println("error here", result.Error.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	newStatus := models.STATUS[requestStatus.Status]
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully update task %d to %s", task.ID, newStatus),
	})
}

func (r *BoardController) DeleteTask(ctx *gin.Context) {
	var task models.Task
	if err := models.DB.First(&task, ctx.Param("task_id")).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Delete(&task).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully delete task %d", task.ID),
	})
}
