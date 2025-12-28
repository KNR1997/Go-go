-- name: ListProducts :many
SELECT
    *
FROM
    products;

-- name: FindProductByID :one
SELECT
    *
FROM
    products
WHERE
    id = $1;

-- name: CreateOrder :one
INSERT INTO orders (
  customer_id
) VALUES ($1) RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items (order_id, product_id, quantity, price_cents)
VALUES ($1, $2, $3, $4) RETURNING *;


-- name: ListCourses :many
SELECT
    *
FROM
    courses;


-- name: CreateCourse :one
INSERT INTO courses (
  name,
  code
) VALUES ($1, $2) RETURNING *;


-- name: FindCourseByID :one
SELECT
    *
FROM
    courses
WHERE
    id = $1;


-- name: UpdateCourse :one
UPDATE courses
SET
  name = $2,
  code = $3
WHERE
  id = $1
RETURNING *;


-- name: DeleteCourse :exec
DELETE FROM courses
WHERE
  id = $1;


-- name: ListQuizzes :many
SELECT
    *
FROM
    quizzes;


-- name: FindQuizByID :one
SELECT
    *
FROM
    quizzes
WHERE
    id = $1;


-- name: CreateQuiz :one
INSERT INTO quizzes (
  course_id,
  week_number,
  date_time,
  status
) VALUES ($1, $2, $3, $4) RETURNING *;


-- name: UpdateQuiz :one
UPDATE quizzes
SET
  week_number = $2,
  date_time = $3,
  status = $4
WHERE
  id = $1
RETURNING *;


-- name: DeleteQuiz :exec
DELETE FROM quizzes
WHERE
  id = $1;


-- name: ListQuizzesWithCourse :many
SELECT
  q.id AS quiz_id,
  q.course_id,
  q.week_number,
  q.date_time,
  q.status,

  c.id AS course_id,
  c.name AS course_name,
  c.code AS course_code
FROM quizzes q
JOIN courses c ON q.course_id = c.id;
