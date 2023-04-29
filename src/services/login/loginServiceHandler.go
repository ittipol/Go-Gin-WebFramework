package login

import (
	"fmt"
	"os"
	"time"
	"web-api/models/request"
	"web-api/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type loginService struct {
	db *gorm.DB
}

type ILoginService interface {
	Login(body *request.LoginBody) (*LoginResponse, error)
	Register(c *gin.Context) error
	// Validate(c *gin.Context)
}

func NewLoginService(db *gorm.DB) ILoginService {
	return &loginService{db}
}

func (h *loginService) Login(body *request.LoginBody) (*LoginResponse, error) {

	var user repository.User

	// check user exist
	result := h.db.Table("users").Where("email = ?", body.Email).Select("id, password").Scan(&user)

	fmt.Println(result.Error)
	fmt.Println(user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	accessTokenSecretKey := []byte(os.Getenv("JWT_ACCESS_TOKEN"))
	refreshTokenSecretKey := []byte(os.Getenv("JWT_REFRESH_TOKEN"))

	// Create the Claims
	claims := MyCustomClaims{
		user.ID,
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
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
		return nil, errAccessToken
	}

	if errRefreshToken != nil {
		return nil, errRefreshToken
	}

	v := &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// b, err := json.Marshal(v)

	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println(string(b))

	return v, nil

}

func (h *loginService) Register(c *gin.Context) error {

	var body request.RegisterBody

	if err := c.ShouldBindJSON(&body); err != nil {
		return err
	}

	password := []byte(body.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := repository.User{
		Email:    body.Email,
		Password: string(hashedPassword),
		Name:     body.Name,
	}

	result := h.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// func (h *loginService) Validate(c *gin.Context) {

// 	// s := c.Request.Header.Get("Authorization")
// 	// tokenString := strings.TrimPrefix(s, "Bearer ")

// 	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOjEsImlzcyI6InRlc3QiLCJleHAiOjE2ODE2Njc3OTF9.Lhz92aJLZ8X7oO5Jyl-P0vWi3M9UP6XvwboxOj276v8"

// 	hmacSampleSecret := []byte("AllYourBase")

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Don't forget to validate the alg is what you expect:
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 		return hmacSampleSecret, nil
// 	})

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		// fmt.Println(claims["foo"], claims["nbf"])

// 		c.JSON(http.StatusOK, gin.H{
// 			"res": claims,
// 		})

// 		return

// 	} else {
// 		fmt.Println(err)

// 		c.JSON(http.StatusForbidden, gin.H{
// 			"error": err.Error(),
// 		})

// 		return
// 	}

// }
