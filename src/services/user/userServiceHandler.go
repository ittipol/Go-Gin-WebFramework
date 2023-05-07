package user

import (
	"errors"
	"web-api/repository"
)

var (
	ErrUserNotFound = errors.New("User Not Found")
)

type UserResponse struct {
	ID    int
	Email string
	Name  string
}

type UserService interface {
	Me(id int) (*UserResponse, error)
}
type userService struct {
	userRepositiry repository.UserRepositiry
}

func NewUserService(userRepositiry repository.UserRepositiry) UserService {
	return &userService{userRepositiry}
}

func (obj *userService) Me(id int) (*UserResponse, error) {

	user, err := obj.userRepositiry.GetUserById(id)

	if err != nil {
		return nil, ErrUserNotFound
	}

	res := UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	return &res, nil

}
