package middleware

import (
	"clean-arsitecture/internal/libraries"
	"os"

	"strings"

	"github.com/gin-gonic/gin"
)

var bearer = "Bearer "

func WithAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(401, gin.H{
				"message": "unauthorized",
			})

			ctx.Abort()
			return
		}

		//untuk cek apakah header = bearer
		if !strings.HasPrefix(authHeader, bearer) {
			ctx.JSON(401, gin.H{

				"message": "unauthorized",
			})

			ctx.Abort()
			return
		}
		token := strings.Split(authHeader, " ")
		secret_key := os.Getenv("JWT_SECRET_KEY")
		data, err := libraries.DecryptJwt(token[1], []byte(secret_key))
		if err != nil {
			ctx.JSON(401, gin.H{
				"message": err.Error(),
			})

			ctx.Abort()
			return
		}
		userID := data["user_id"]
		role := data["role"]
		name := data["name"]
		ctx.Set("user_id", userID)
		ctx.Set("role", role)
		ctx.Set("name", name)

		ctx.Next()
	}

}
