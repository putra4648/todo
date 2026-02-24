package main

import (
	"context"
	"putra4648/todo/db"
	"putra4648/todo/handlers"
	"putra4648/todo/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {

	// load env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := gin.Default()

	app.Use(middleware.Logger())
	app.Use(middleware.Error())
	app.Use(middleware.Auth())

	// app.LoadHTMLGlob("templates/**/*")

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=postgres password=admin123 host=192.168.1.6 port=5432 dbname=todo sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

	db := db.New(conn)

	auth := app.Group("/auth")
	auth.POST("/register", handlers.RegisterHandler(db))
	auth.POST("/login", handlers.LoginHandler(db))

	app.GET("/todos", handlers.GetTodosHandler(db))
	app.POST("/todos", handlers.CreateTodoHandler(db))
	app.PUT("/todos/:id", handlers.UpdateTodoHandler(db))
	app.DELETE("/todos/:id", handlers.DeleteTodoHandler(db))

	app.Run(":8080")
}
