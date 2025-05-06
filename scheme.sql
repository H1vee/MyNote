CREATE DATABASE MyNote;
CREATE TABLE users(
            id SERIAL PRIMARY KEY,
            email TEXT NOT NULL,
            hash_password TEXT NOT NULL
                  );

CREATE TABLE notes(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL
);