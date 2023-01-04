package service

import (
	"clean-arsitecture/internal/domain"
	"clean-arsitecture/internal/libraries"
	"errors"
	"os"
)

type AuthService struct {
	authRepo domain.AuthRepository
}

func NewAuthService(ar domain.AuthRepository) *AuthService {
	return &AuthService{
		authRepo: ar,
	}
}

func (as *AuthService) Login(login domain.Login) (string, error, int) {
	//cek email
	res, err := as.authRepo.GetUserByEmail(login)
	if err != nil {
		return "", errors.New("Failed Email or password"), 401
	}

	//cek password
	err = libraries.ComparePassword(res.Password, login.Password)
	if err != nil {
		return "", errors.New("Failed Email or password"), 401
	}

	//set token jwt
	secret_key := os.Getenv("JWT_SECRET_KEY")
	jwt, err := libraries.GenerateJWT(res.ID, res.RoleId, res.Name, []byte(secret_key))
	if err != nil {
		return "", domain.ErrInternalServerError, 500
	}

	return jwt, nil, 200

}
