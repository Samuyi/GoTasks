CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL,
    title text NOT NULL,
    done BOOLEAN DEFAULT FALSE,
    user_id INTEGER REFERENCES Users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)

CREATE TABLE IF NOT EXISTS Users (
id  SERIAL, 
name text NOT NULL,
password text NOT NULL,
email text NOT NULL UNIQUE,
PRIMARY KEY(id)
)