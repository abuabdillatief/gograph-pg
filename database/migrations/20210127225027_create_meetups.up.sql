CREATE TABLE meetups(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    -- meaning that user_id will refer to the column id of table users
    user_id BIGSERIAL REFERENCES users(id) ON DELETE CASCADE NOT NULL
)