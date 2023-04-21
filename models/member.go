package models

import "gorm.io/gorm"

var ROLES map[int]string = map[int]string{
	1: "ADMIN",
	2: "MEMBER",
}

var MEMBER_ADMIN int = 1
var MEMBER_MEMBER int = 2

type Member struct {
	gorm.Model
	UserId  int `form:"user_id" json:"user_id" binding:"required"`
	BoardId int `form:"board_id" json:"board_id" binding:"required"`
	Role    int `form:"role" json:"role" binding:"required"`
	User    User
}

func (m *Member) GetRole() string {
	return ROLES[m.Role]
}
