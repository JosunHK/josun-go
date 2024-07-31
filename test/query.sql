-- name: GetUsers :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users 
ORDER BY name;

-- name: CreateUsers :exec
INSERT INTO users(
  name, email, password 
) VALUES (
  $1, $2, $3
);

-- name: UpdateUsers :exec
UPDATE users
  set name = $2,
  email = $3,
  password = $3
WHERE id = $1;

-- name: DeleteUsers :exec
DELETE FROM users 
WHERE id = $1;
