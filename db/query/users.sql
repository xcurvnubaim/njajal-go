-- name: CreateUsers :exec
INSERT INTO users (
    email,
    password
) VALUES (
    $1, $2
) RETURNING *;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: ResetPassword :exec
UPDATE users SET password = $1 WHERE email = $2;