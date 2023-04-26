package models

import "gorm.io/gorm"

var STATUS map[int]string = map[int]string{
	1: "TODO",
	2: "IN_PROGRESS",
	3: "REVIEW",
	4: "DONE",
}

var (
	STATUS_TODO        int = 1
	STATUS_IN_PROGRESS int = 2
	STATUS_REVIEW      int = 3
	STATUS_DONE        int = 4
)

type Task struct {
	gorm.Model
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	BoardId     int    `form:"board_id" json:"board_id" binding:"required"`
	Status      int    `form:"status" json:"status"`
	AssigneeId  int    `form:"assignee_id" json:"assingee_id"`
	Assignee    User   `form:"assignee" json:"assignee"`
}

type TaskResponse struct {
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	BoardId     int         `json:"board_id"`
	Status      int         `json:"status"`
	AssigneId   int         `json:"assignee_id"`
	Assignee    UserReponse `json:"assignee"`
}

type TaskRequest struct {
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	BoardId     int    `form:"board_id" json:"board_id" binding:"required"`
	// AssigneeId  int    `form:"assignee_id" json:"assignee_id" binding:"required"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.Status = STATUS_TODO
	return
}

func (t *Task) GetStatus() string {
	return STATUS[t.Status]
}

func (t *Task) Unassigned() bool {
	return t.AssigneeId == 0
}

func (t *Task) GetReponse() TaskResponse {
	return TaskResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		BoardId:     t.BoardId,
		Status:      t.Status,
		AssigneId:   t.AssigneeId,
		Assignee:    t.Assignee.GetResponse(),
	}
}
