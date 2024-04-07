-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, price, category_id)
VALUES (?, ?, ?, ?, ?);
