package models

import (
	"errors"

	"github.com/aldinofrizal/bumpaban/helpers"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserReponse struct {
	ID    int    `form:"id" json:"id"`
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
}

type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPw, err := helpers.HashPassword(u.Password)
	if err != nil {
		err = errors.New(err.Error())
	}
	u.Password = hashedPw
	return err
}

func (u *User) GetResponse() UserReponse {
	return UserReponse{
		ID:    int(u.ID),
		Name:  u.Name,
		Email: u.Email,
	}
}
