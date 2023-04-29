package main

import (
	"fmt"
	"os"
	"web-api/controllers"
	"web-api/middlewares"
	"web-api/orm/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		panic("Cannot load .env")
	}

	fmt.Println(os.Environ())

	mode, found := os.LookupEnv("ENV")

	fmt.Println("ENV: ", os.Getenv("ENV"))

	if found && mode == "production" {
		fmt.Println("Run on production mode")
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	fmt.Println("Start....")
	// mode := os.Getenv(EnvGinMode)

	// init DB
	db := db.GetConnection(os.Getenv("DSN"))
	// defer db.close()

	// controllers
	var healthController controllers.IHealthController = controllers.NewHealthController()
	var loginController controllers.ILoginController = controllers.NewLoginController(db)
	var userController controllers.IUserController = controllers.NewUserController(db)

	r := gin.Default()

	user := r.Group("user")
	user.Use(middlewares.AuthorizeJWT)
	{
		user.GET("/profile", userController.Me)
	}

	r.GET("/health", healthController.Health)
	// r.Use(middlewares.CheckAuth).GET("/login", loginController.Login)
	r.POST("/login", loginController.Login)
	r.POST("/register", loginController.Register)
	// r.GET("/validate", services.Validate)

	port, found := os.LookupEnv("PORT")

	if !found {
		panic("Port does not specific")
	}

	if err := r.Run(":" + port); err != nil {
		panic("Cannot run on port " + port)
	}
}
