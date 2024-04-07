-- name: CreateCourse :exec
INSERT INTO courses (id, name, category_id) VALUES (?, ?, ?);

