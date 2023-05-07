package login_test

import (
	"errors"
	"testing"
	"web-api/repository"
	"web-api/services/login"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	userRepositoryMock := repository.NewUserRepositoryMock()

	// Arrange
	userRepositoryMock.On("GetUserByEmail", "email").Return(&repository.User{
		ID:       1,
		Email:    "m1@email.com",
		Password: "$2a$12$ErV8S.A18n1nHftm5ph1Z./KhW/Kx41IVrliRSf1XuInN6TyPIAtC",
		Name:     "Mock 1",
	}, nil)

	userRepositoryMock.On("GetUserByEmail", "email2").Return(&repository.User{}, errors.New("Not Found"))

	userRepositoryMock.On("GetUserByEmail", "email3").Return(&repository.User{
		ID:       3,
		Email:    "m3@email.com",
		Password: "$2a$12$ErV8S.A18n1nHftm5ph1Z./KhW/Kx41IVrliRSf1XuInN6TyPIAtC",
		Name:     "Mock 3",
	}, nil)

	// Act
	t.Run("Login success", func(t *testing.T) {
		svc := login.NewLoginService(userRepositoryMock)

		res, err := svc.Login("email", "1234")

		assert.Equal(t, nil, err)

		assert.NotEmpty(t, res.AccessToken)
		assert.NotEmpty(t, res.RefreshToken)
	})

	t.Run("Login user not found", func(t *testing.T) {

		svc := login.NewLoginService(userRepositoryMock)

		res, err := svc.Login("email2", "1234")

		assert.Empty(t, res)
		assert.EqualError(t, err, "Email or password is invalid")
	})

	t.Run("Login password not match", func(t *testing.T) {

		svc := login.NewLoginService(userRepositoryMock)

		res, err := svc.Login("email3", "pw")

		assert.Empty(t, res)
		assert.EqualError(t, err, "Email or password is invalid")
	})
}

func TestRegister(t *testing.T) {
	userRepositoryMock := repository.NewUserRepositoryMock()

	userRepositoryMock.On("CreateUser", "email", "Mock 1").Return(&repository.User{
		ID:       1,
		Email:    "m1@email.com",
		Password: "$2a$12$ErV8S.A18n1nHftm5ph1Z./KhW/Kx41IVrliRSf1XuInN6TyPIAtC",
		Name:     "Mock 1",
	}, nil)

	userRepositoryMock.On("CreateUser", "email2", "Mock 2").Return(&repository.User{}, errors.New("Cannot Create User"))

	t.Run("Register success", func(t *testing.T) {

		svc := login.NewLoginService(userRepositoryMock)

		err := svc.Register("email", "1234", "Mock 1")

		assert.Empty(t, err)
	})

	t.Run("Register email empty", func(t *testing.T) {

		svc := login.NewLoginService(userRepositoryMock)

		err := svc.Register("", "1234", "Mock 1")

		assert.EqualError(t, err, "Cannot Create User")
	})

	t.Run("Register password empty", func(t *testing.T) {

		svc := login.NewLoginService(userRepositoryMock)

		err := svc.Register("email", "", "Mock 1")

		assert.EqualError(t, err, "Cannot Create User")
	})

	t.Run("Register name empty", func(t *testing.T) {

		svc := login.NewLoginService(userRepositoryMock)

		err := svc.Register("email", "1234", "")

		assert.EqualError(t, err, "Cannot Create User")
	})

	t.Run("Register create user fail", func(t *testing.T) {

		svc := login.NewLoginService(userRepositoryMock)

		err := svc.Register("email2", "1234", "Mock 2")

		assert.EqualError(t, err, "Cannot Create User")
	})
}

func BenchmarkLogin(b *testing.B) {

	userRepositoryMock := repository.NewUserRepositoryMock()

	userRepositoryMock.On("Login", "email").Return(&repository.User{
		ID:       1,
		Email:    "m1@email.com",
		Password: "$2a$12$ErV8S.A18n1nHftm5ph1Z./KhW/Kx41IVrliRSf1XuInN6TyPIAtC",
		Name:     "Mock 1",
	}, nil)

	svc := login.NewLoginService(userRepositoryMock)

	for i := 0; i < b.N; i++ {
		svc.Login("email", "1234")
	}
}
