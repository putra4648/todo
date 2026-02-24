CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_todos (
    user_id INTEGER NOT NULL,
    todo_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, todo_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (todo_id) REFERENCES todos(id)
);
