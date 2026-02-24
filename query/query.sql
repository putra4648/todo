-- name: CreateUser :one
INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: CreateTodoWithUser :one
WITH inserted_todo AS (
    INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id
)
INSERT INTO user_todos (user_id, todo_id)
SELECT $4, id FROM inserted_todo
RETURNING todo_id;

-- name: GetTodo :one
SELECT * FROM todos WHERE id = $1;

-- name: GetTodosByUserID :many
SELECT t.id, t.title, t.description, t.completed, t.created_at, t.updated_at FROM todos t
JOIN user_todos ut ON t.id = ut.todo_id
WHERE ut.user_id = $1;

-- name: UpdateTodo :one
UPDATE todos
SET title = $1, description = $2, completed = $3
WHERE id = $4 AND id IN (SELECT todo_id FROM user_todos WHERE user_id = $5)
RETURNING id;

-- name: DeleteTodo :one
DELETE FROM todos
WHERE id = $1 AND id IN (SELECT todo_id FROM user_todos WHERE user_id = $2)
RETURNING id;

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