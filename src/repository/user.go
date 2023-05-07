package repository

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	ID           int    `gorm:"primaryKey:autoIncrement"`
	Email        string `gorm:"unique"`
	Password     string
	Name         string
	RefreshToken string
}

type UserRepositiry interface {
	GetUserById(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(email string, hashedPassword string, name string) (*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepositiry(db *gorm.DB) UserRepositiry {
	return &userRepository{db}
}

func (obj *userRepository) GetUserById(id int) (*User, error) {
	user := User{}

	tx := obj.db.Table("users").Where("id = @id", sql.Named("id", id)).Select("id, email, name").Scan(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (obj *userRepository) GetUserByEmail(email string) (*User, error) {

	user := User{}

	tx := obj.db.Table("users").Where("email = @email", sql.Named("email", email)).Select("id, password").Scan(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func (obj *userRepository) CreateUser(email string, hashedPassword string, name string) (*User, error) {

	user := User{
		Email:    name,
		Password: hashedPassword,
		Name:     name,
	}

	tx := obj.db.Create(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}
