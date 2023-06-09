package models

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Title   string   `form:"title" json:"title" binding:"required"`
	Members []Member `json:"members"`
	Users   []User   `gorm:"many2many:members" json:"users"`
	Tasks   []Task   `json:"tasks"`
}

type BoardResponse struct {
	ID    uint                  `json:"id"`
	Title string                `json:"title"`
	Users []UserReponseWithRole `json:"members"`
	Tasks []TaskResponse        `json:"tasks"`
}

func (b *Board) GetOwnerId() int {
	for _, member := range b.Members {
		if member.Role == MEMBER_ADMIN {
			return member.UserId
		}
	}
	return 0
}

func (b *Board) HasMember(user_id int) bool {
	for _, member := range b.Members {
		if member.UserId == user_id {
			return true
		}
	}
	return false
}

func (b *Board) GetIndexResponse() BoardResponse {
	members := []UserReponseWithRole{}
	for _, member := range b.Members {
		members = append(members, UserReponseWithRole{
			ID:    int(member.User.ID),
			Name:  member.User.Name,
			Email: member.User.Email,
			Role:  member.GetRole(),
		})
	}

	tasks := []TaskResponse{}
	for _, task := range b.Tasks {
		tasks = append(tasks, task.GetReponse())
	}

	return BoardResponse{
		ID:    b.ID,
		Title: b.Title,
		Users: members,
		Tasks: tasks,
	}
}
