-- name: CreateUser :one
INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateTodo :one
INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id;

-- name: GetTodo :one
SELECT * FROM todos WHERE id = $1;

-- name: GetTodos :many
SELECT * FROM todos;

-- name: UpdateTodo :one
UPDATE todos SET title = $1, description = $2, completed = $3 WHERE id = $4 RETURNING id;

-- name: DeleteTodo :one
DELETE FROM todos WHERE id = $1 RETURNING id;

-- name: CreateUserTodo :one
INSERT INTO user_todos (user_id, todo_id) VALUES ($1, $2) RETURNING todo_id;

-- name: GetUserTodo :one
SELECT * FROM user_todos WHERE user_id = $1 AND todo_id = $2;

-- name: GetUserTodos :many
SELECT * FROM user_todos WHERE user_id = $1;

-- name: GetTodoUsers :many
SELECT * FROM user_todos WHERE todo_id = $1;

-- name: DeleteUserTodo :one
DELETE FROM user_todos WHERE user_id = $1 AND todo_id = $2 RETURNING todo_id;