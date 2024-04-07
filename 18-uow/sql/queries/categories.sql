-- name: CreateCategory :exec
INSERT INTO categories (id, name) VALUES (?, ?);
