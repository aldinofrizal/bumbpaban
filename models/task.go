package models

import "gorm.io/gorm"

var STATUS map[int]string = map[int]string{
	1: "TODO",
	2: "IN_PROGRESS",
	3: "DONE",
}

var (
	STATUS_TODO        int = 1
	STATUS_IN_PROGRESS int = 2
	STATUS_DONE        int = 3
)

type Task struct {
	gorm.Model
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	BoardId     int    `form:"board_id" json:"board_id" binding:"required"`
	Status      int    `form:"status" json:"status"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.Status = STATUS_TODO
	return
}

func (t *Task) GetStatus() string {
	return STATUS[t.Status]
}
