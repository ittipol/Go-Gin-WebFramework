package repository

import "github.com/stretchr/testify/mock"

type userRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *userRepositoryMock {
	return &userRepositoryMock{}
}

func (m *userRepositoryMock) GetUserById(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *userRepositoryMock) GetUserByEmail(email string) (*User, error) {
	args := m.Called(email)
	return args.Get(0).(*User), args.Error(1)
}

func (m *userRepositoryMock) CreateUser(email string, hashedPassword string, name string) (*User, error) {
	args := m.Called(email, name)
	return args.Get(0).(*User), args.Error(1)
}
