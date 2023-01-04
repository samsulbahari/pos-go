package handler

import (
	"clean-arsitecture/internal/domain"
	"clean-arsitecture/internal/libraries"
	"clean-arsitecture/internal/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	Message string `json:"message"`
}

type AuthHandler struct {
	authservice domain.AuthService
}

func NewAuthHandler(r *gin.Engine, as domain.AuthService) {
	handler := &AuthHandler{
		authservice: as,
	}
	auth := r.Group("auth")
	{
		auth.POST("login", handler.Login)
		auth.Use(middleware.WithAuth())
		auth.POST("refresh_token", handler.RefreshToken)

	}

}
func (ah *AuthHandler) Login(ctx *gin.Context) {
	var login domain.Login
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		validation_response := libraries.Validation(err)
		ctx.JSON(422, gin.H{
			"message": validation_response,
		})
		return
	}

	res, err, code := ah.authservice.Login(login)
	if err != nil {
		ctx.JSON(code, ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(code, gin.H{
		"message": "success Login",
		"token":   res,
	})

}

func (ah *AuthHandler) RefreshToken(ctx *gin.Context) {
	roleID := ctx.MustGet("role").(float64)
	userID := ctx.MustGet("user_id").(float64)
	name := ctx.MustGet("name")
	fmt.Println(roleID, userID, name)
	secret_key := os.Getenv("JWT_SECRET_KEY")
	token, err := libraries.GenerateJWT(int(userID), int(roleID), name.(string), []byte(secret_key))
	if err != nil {
		ctx.JSON(401, ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "success refresh token",
		"token":   token,
	})
}
