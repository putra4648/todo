package handlers

import (
	"net/http"
	"putra4648/todo/db"

	"github.com/gin-gonic/gin"
)

func GetTodosHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Todo handler"})
	}
}
