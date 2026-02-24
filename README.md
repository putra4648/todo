# Go Todo API

A robust and secure Todo List API built with Go, Gin, and PostgreSQL. This project follows clean architectural patterns, uses `sqlc` for type-safe database interactions, and implements JWT-based authentication with dual support for Cookies and Authorization headers.

## 🚀 Features

- **User Authentication**: Secure Registration and Login using JWT.
- **Dual Auth Support**: Works with both `Authorization: Bearer <token>` headers and HTTP-only Cookies.
- **User Isolation**: Users can only create, view, update, and delete their own todos.
- **Type-Safe Database**: Uses `sqlc` for generating type-safe Go code from SQL queries.
- **Centralized Error Handling**: Custom middleware for consistent JSON error responses.
- **Logging**: Integrated request logging middleware.

## 🛠️ Tech Stack

- **Languange**: [Go (Golang)](https://golang.org/)
- **Web Framework**: [Gin Gonic](https://gin-gonic.com/)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **ORM/Query Builder**: [sqlc](https://sqlc.dev/)
- **Driver**: [pgx/v5](https://github.com/jackc/pgx)
- **Security**: JWT (golang-jwt/v5) & bcrypt for password hashing.

## 📁 Project Structure

```text
├── db/              # Generated sqlc code (DB interfaces and models)
├── handlers/        # API route handlers (Auth and Todo logic)
├── middleware/      # Gin middlewares (Auth, Log, Error)
├── migrations/      # SQL migration files
├── models/          # Data Transfer Objects (DTOs) for requests/responses
├── query/           # Raw SQL queries used by sqlc
├── utils/           # Helper functions (Hashing, etc.)
├── main.go          # Application entry point
├── sqlc.yaml        # sqlc configuration
└── .env             # Environment variables (not tracked)
```

## ⚙️ Setup & Installation

### 1. Prerequisites

- Go 1.21+
- PostgreSQL
- [sqlc](https://sqlc.dev/) (optional, for regenerating queries)

### 2. Database Setup

Create a PostgreSQL database and apply the migrations. You can find the schema in the `query/` or `migrations/` directory.

### 3. Environment Variables

Create a `.env` file in the root directory:

```env
JWT_SECRET=your_secret_key_here
```

### 4. Running the Application

```bash
go mod tidy
go run main.go
```

The server will start at `http://localhost:8080`.

## 🛣️ API Endpoints

### Authentication

| Method | Endpoint         | Description                           |
| :----- | :--------------- | :------------------------------------ |
| POST   | `/auth/register` | Register a new user                   |
| POST   | `/auth/login`    | Login and receive JWT (Cookie & JSON) |

### Todos (Requires Auth)

| Method | Endpoint     | Description                           |
| :----- | :----------- | :------------------------------------ |
| GET    | `/todos`     | List all todos for the logged-in user |
| POST   | `/todos`     | Create a new todo                     |
| PUT    | `/todos/:id` | Update an existing todo               |
| DELETE | `/todos/:id` | Delete a todo                         |

## 🛠️ Development

### Generating Database Code

If you modify `query/query.sql`, regenerate the Go code using:

```bash
sqlc generate
```

## 📜 License

This project is open-source and available under the MIT License.
