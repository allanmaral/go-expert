-- name: CreateCourse :exec
INSERT INTO courses (id, name, description, price, category_id)
VALUES (?, ?, ?, ?, ?);

-- name: ListCourses :many
SELECT c.*,
       ca.name AS category_name
FROM courses c
    JOIN categories ca ON c.category_id = ca.id;
