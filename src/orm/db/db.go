package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Init(dsn string) {
	Conn = GetConnection(dsn)
}

func GetConnection(dsn string) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DryRun: false,
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database Connected")

	return conn
}
