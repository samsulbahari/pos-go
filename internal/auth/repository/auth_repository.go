package repository

import (
	"clean-arsitecture/internal/domain"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (m *AuthRepository) GetUserByEmail(login domain.Login) (domain.User, error) {
	var user domain.User
	err := m.db.Where("email = ?", login.Email).First(&user).Error
	return user, err
}
