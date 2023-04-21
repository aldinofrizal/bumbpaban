package services

import (
	"errors"

	"github.com/aldinofrizal/bumpaban/models"
)

func CreateBoard(board *models.Board, owner *models.User) error {
	tx := models.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(board).Error; err != nil {
		tx.Rollback()
		return err
	}
	member := models.Member{
		UserId:  int(owner.ID),
		BoardId: int(board.ID),
		Role:    models.MEMBER_ADMIN,
	}
	if err := tx.Create(&member).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func AddMember(board *models.Board, user_id int) error {
	tx := models.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	var user models.User
	if err := tx.First(&user, user_id).Error; err != nil {
		tx.Rollback()
		return err
	}

	member := models.Member{}
	result := tx.Where("user_id = ? AND board_id = ?", user.ID, board.ID).First(&member)
	if result.RowsAffected > 0 {
		tx.Rollback()
		return errors.New("already member")
	}

	member.UserId = int(user.ID)
	member.BoardId = int(board.ID)
	member.Role = models.MEMBER_MEMBER

	if err := tx.Create(&member).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func FormatBoards(boards []models.Board) []models.BoardResponse {
	result := []models.BoardResponse{}
	for _, board := range boards {
		result = append(result, board.GetIndexResponse())
	}

	return result
}
