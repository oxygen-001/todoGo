CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    done BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);