package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.FullPath() == "/auth/register" || c.FullPath() == "/auth/login" {
			c.Next()
			return
		}

		authorizationHeader := c.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := jwt.ParseWithClaims(authorizationHeader, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// check expired time
		if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
			exp, err := claims.GetExpirationTime()
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
			if exp.Before(time.Now()) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
				return
			}
		}

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			return
		}

	}
}
