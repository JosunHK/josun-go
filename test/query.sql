-- name: GetUser :one
SELECT * FROM users 
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users 
ORDER BY name;

-- name: CreateUser :execresult
INSERT INTO users(
  name, email, password 
) VALUES (
  ?, ?, ?
);

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = ?;
