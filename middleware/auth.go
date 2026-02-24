package middleware

import (
	"errors"
	"os"
	"strings"
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

		tokenString := c.Request.Header.Get("Authorization")
		if tokenString != "" {
			tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		}

		if tokenString == "" {
			cookie, err := c.Cookie("token")
			if err == nil {
				tokenString = cookie
			}
		}

		if tokenString == "" {
			c.Error(errors.New("Unauthorized"))
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		if !token.Valid {
			c.Error(errors.New("Token is not valid"))
			c.Abort()
			return
		}

		// check expired time and set user_id
		if claims, ok := token.Claims.(*jwt.MapClaims); ok {
			exp, err := claims.GetExpirationTime()
			if err != nil {
				c.Error(err)
				c.Abort()
				return
			}
			if exp != nil && exp.Before(time.Now()) {
				c.Error(errors.New("Token is expired"))
				c.Abort()
				return
			}

			if userID, ok := (*claims)["user_id"].(float64); ok {
				c.Set("user_id", int32(userID))
			}
		}

		c.Next()
	}
}
