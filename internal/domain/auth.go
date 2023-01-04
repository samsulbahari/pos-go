package domain

import "time"

type User struct {
	ID         int
	Email      string
	Name       string
	RoleId     int
	Password   string
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthRepository interface {
	GetUserByEmail(login Login) (User, error)
}

type AuthService interface {
	Login(login Login) (string, error, int)
}
