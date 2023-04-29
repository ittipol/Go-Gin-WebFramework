package user

import (
	"web-api/repository"

	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

type IUserService interface {
	Me(id int) (*repository.User, error)
}

func NewUserService(db *gorm.DB) IUserService {
	return &userService{
		db,
	}
}

func (h *userService) Me(id int) (*repository.User, error) {

	user := repository.User{}

	result := h.db.Table("users").Where("id = ?", id).Select("id, email, name").Scan(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}
