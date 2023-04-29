package repository

type User struct {
	ID           int    `json:"id";db:"id";gorm:"primaryKey:autoIncrement"`
	Email        string `json:"email";db:"email";gorm:"unique"`
	Password     string `json:"password";db:"password"`
	Name         string `json:"name";db:"name"`
	RefreshToken string `json:"refresh_token";db:"refresh_token"`
}

type IUserRepositiry interface {
}

type userRepository struct {
}

func UserRepository() IUserRepositiry {
	return &userRepository{}
}
