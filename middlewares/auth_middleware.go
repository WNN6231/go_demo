package middlewares

import (
	"exchangeapp/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(401, gin.H{"error": "Missing authorization token"})
			ctx.Abort()
			return
		}

		username, err := utils.ParseJWT(token)
		if err != nil {
			ctx.JSON(401, gin.H{"error": "Invalid token: " + err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("username", username)
		ctx.Next()
	}
}
