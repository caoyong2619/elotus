package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/caoyong2619/elotus/internal/services"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestToken := c.Request.Header.Get("Authorization")
		if requestToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 1,
				"msg":  "unauthorized",
				"data": nil,
			})
			c.Abort()
			return
		}

		//var token *jwt.Token
		token, err := authService.ParseToken(requestToken)

		if err != nil {
			// record error
			slog.Error("parse token failed", "err", err)

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 1,
				"msg":  err.Error(),
				"data": nil,
			})
			c.Abort()
			return
		}

		c.Set("token", token.Claims.(*services.ElotusClaims))
	}
}
