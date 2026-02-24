package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		time := time.Now()
		c.Next()
		fmt.Println(time, c.Request.Method, c.Request.URL, c.Writer.Status())
	}
}
