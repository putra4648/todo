package handlers

import (
	"errors"
	"net/http"
	"os"
	"putra4648/todo/db"
	"putra4648/todo/models"
	"putra4648/todo/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.AuthRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			return
		}
		pw, err := utils.HashPassword(req.Password)

		if err != nil {
			c.Error(err)
			return
		}

		q.CreateUser(c.Request.Context(), db.CreateUserParams{
			Name:     req.Username,
			Password: pw,
		})

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}

func LoginHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.AuthRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(err)
			return
		}

		user, err := q.GetUserByName(c.Request.Context(), req.Username)
		if err != nil {
			c.Error(errors.New("User not found"))
			return
		}

		if utils.VerifyPassword(req.Password, user.Password) != nil {
			c.Error(errors.New("Password not match"))
			return
		}

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
		}).SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			c.Error(err)
			return
		}

		c.SetCookie("token", token, 60*60*24*7, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "User logged in successfully",
			"token":   token,
		})
	}
}
