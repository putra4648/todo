package middleware

import (
	"os"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		var userID int32
		if tokenString != "" {
			token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if err == nil && token.Valid {
				if claims, ok := token.Claims.(*jwt.MapClaims); ok {
					if id, ok := (*claims)["user_id"].(float64); ok {
						userID = int32(id)
						c.Set("user_id", userID)
						c.Set("is_logged_in", true)
					}
				}
			}
		}

		path := c.FullPath()
		actualPath := c.Request.URL.Path
		allowedUrls := []string{
			"/auth/register",
			"/auth/login",
			"/",
			"/login",
			"/register",
			"/favicon.ico",
		}

		isAllowed := slices.Contains(allowedUrls, path) || slices.Contains(allowedUrls, actualPath)

		// Redirect logged in users away from login/register
		if userID != 0 && (actualPath == "/login" || actualPath == "/register") {
			c.Redirect(302, "/")
			c.Abort()
			return
		}

		if !isAllowed && userID == 0 {
			c.Error(NewAppError(401, "Unauthorized"))
			c.Abort()
			return
		}

		c.Next()
	}
}
