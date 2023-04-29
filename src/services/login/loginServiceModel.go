package login

import "github.com/golang-jwt/jwt/v5"

type MyCustomClaims struct {
	Id int `json:"userId"`
	jwt.RegisteredClaims
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
