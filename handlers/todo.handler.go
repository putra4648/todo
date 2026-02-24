package handlers

import (
	"net/http"
	"putra4648/todo/db"
	"putra4648/todo/middleware"
	"putra4648/todo/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func GetTodosHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.Error(middleware.NewAppError(401, "Unauthorized"))
			return
		}

		todos, err := q.GetTodosByUserID(c.Request.Context(), userID.(int32))
		if err != nil {
			c.Error(middleware.NewAppError(500, err.Error()))
			return
		}

		var todoDto []models.TodoDto = make([]models.TodoDto, 0)
		for _, todo := range todos {
			todoDto = append(todoDto, models.TodoDto{
				ID:          todo.ID,
				Title:       todo.Title,
				Description: todo.Description.String,
				IsDone:      todo.Completed.Bool,
				CreatedAt:   todo.CreatedAt.Time.String(),
				UpdatedAt:   todo.UpdatedAt.Time.String(),
			})
		}
		c.JSON(http.StatusOK, todoDto)
	}
}

func CreateTodoHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.Error(middleware.NewAppError(401, "Unauthorized"))
			return
		}

		var req models.TodoDto
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(middleware.NewAppError(400, err.Error()))
			return
		}

		q.CreateTodoWithUser(c.Request.Context(), db.CreateTodoWithUserParams{
			Title:       req.Title,
			Description: pgtype.Text{String: req.Description, Valid: true},
			Completed:   pgtype.Bool{Bool: req.IsDone, Valid: true},
			UserID:      userID.(int32),
		})
		c.JSON(http.StatusOK, gin.H{"message": "Todo created successfully"})
	}
}

func UpdateTodoHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.Error(middleware.NewAppError(401, "Unauthorized"))
			return
		}

		idParam := c.Param("id")
		parseId, err := strconv.Atoi(idParam)
		if err != nil {
			c.Error(middleware.NewAppError(400, "Invalid ID format"))
			return
		}

		var req models.TodoDto
		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(middleware.NewAppError(400, err.Error()))
			return
		}

		id, err := q.UpdateTodo(c.Request.Context(), db.UpdateTodoParams{
			ID:          int32(parseId),
			Title:       req.Title,
			Description: pgtype.Text{String: req.Description, Valid: true},
			Completed:   pgtype.Bool{Bool: req.IsDone, Valid: true},
			UserID:      userID.(int32),
		})

		if err != nil {
			c.Error(middleware.NewAppError(500, err.Error()))
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "id": id})
	}
}

func DeleteTodoHandler(q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.Error(middleware.NewAppError(401, "Unauthorized"))
			return
		}

		id := c.Param("id")

		parseId, err := strconv.Atoi(id)
		if err != nil {
			c.Error(middleware.NewAppError(400, "Invalid ID format"))
			return
		}
		q.DeleteTodo(c.Request.Context(), db.DeleteTodoParams{
			ID:     int32(parseId),
			UserID: userID.(int32),
		})
		c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
	}
}
