package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			// Step3: Use the last error
			err := c.Errors.Last().Err

			// Step4: Respond with a generic error message
			c.JSON(http.StatusInternalServerError, map[string]any{
				"success": false,
				"message": err.Error(),
			})
		}
	}
}
