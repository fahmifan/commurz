-- name: FindCartByUserID :one
SELECT * FROM carts WHERE user_id = @user_id;

-- name: FindUserByID :one
SELECT * FROM users WHERE id = @id;

-- name: CreateUser :one
INSERT INTO users (id, email)
VALUES (@id, @email)
RETURNING *;

-- name: FindAllUsers :many
SELECT * FROM users;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = @email LIMIT 1;
