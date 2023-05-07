package login

import (
	"errors"
	"os"
	"time"
	"web-api/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredential = errors.New("Email or password is invalid")
	ErrUserNotFound      = errors.New("User Not Found")
	ErrCannotCreateUser  = errors.New("Cannot Create User")
)

type LoginService interface {
	Login(email string, password string) (*LoginResponse, error)
	Register(email string, password string, name string) error
}
type loginService struct {
	userRepositiry repository.UserRepositiry
}

func NewLoginService(userRepositiry repository.UserRepositiry) LoginService {
	return &loginService{userRepositiry}
}

func (obj *loginService) Login(email string, password string) (*LoginResponse, error) {

	// check user exist
	user, err := obj.userRepositiry.GetUserByEmail(email)

	if err != nil {
		return nil, ErrInvalidCredential
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, ErrInvalidCredential
	}

	accessTokenSecretKey := []byte(os.Getenv("JWT_ACCESS_TOKEN"))
	refreshTokenSecretKey := []byte(os.Getenv("JWT_REFRESH_TOKEN"))

	// Create the Claims
	claims := MyCustomClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(5*time.Minute).Unix(), 0)),
		},
	}

	// Access Token
	_accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Refresh token
	_refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(30*time.Minute).Unix(), 0)),
	})

	accessToken, errAccessToken := _accessToken.SignedString(accessTokenSecretKey)
	refreshToken, errRefreshToken := _refreshToken.SignedString(refreshTokenSecretKey)

	if errAccessToken != nil {
		return nil, ErrInvalidCredential
	}

	if errRefreshToken != nil {
		return nil, ErrInvalidCredential
	}

	res := &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return res, nil
}

func (obj *loginService) Register(email string, password string, name string) error {

	if email == "" || password == "" || name == "" {
		return ErrCannotCreateUser
	}

	passwordByte := []byte(password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = obj.userRepositiry.CreateUser(email, string(hashedPassword), name)

	if err != nil {
		return ErrCannotCreateUser
	}

	return nil
}
