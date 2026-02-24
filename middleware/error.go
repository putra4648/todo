package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	statusCode uint
	message    string
}

func (e *AppError) Error() string {
	return e.message
}

func (e *AppError) StatusCode() uint {
	return e.statusCode
}

func NewAppError(status uint, message string) *AppError {
	return &AppError{
		statusCode: status,
		message:    message,
	}
}

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			if c.Writer.Written() {
				return
			}

			// Step3: Use the last error
			err := c.Errors.Last().Err

			// Step4: Respond with a generic error message
			var appError *AppError
			if ok := errors.As(err, &appError); ok {
				c.JSON(int(appError.StatusCode()), map[string]any{
					"success": false,
					"message": appError.Error(),
				})
			} else {
				c.JSON(http.StatusInternalServerError, map[string]any{
					"success": false,
					"message": err.Error(),
				})
			}
		}
	}
}
