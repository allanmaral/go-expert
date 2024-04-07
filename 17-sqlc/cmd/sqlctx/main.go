package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/allanmaral/go-expert/17-sqlc/internal/db"
	"github.com/google/uuid"
	// "github.com/google/uuid"

	_ "github.com/go-sql-driver/mysql"
)

type CourseAR struct {
	db *sql.DB
	*db.Queries
}

func NewCourseAR(dbconn *sql.DB) *CourseAR {
	return &CourseAR{
		db:      dbconn,
		Queries: db.New(dbconn),
	}
}

func (c *CourseAR) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("error on rollback: %v, original error %w", rbErr, err)
		}
		return err
	}

	return tx.Commit()
}

func (c *CourseAR) CreateCourseAndCategory(ctx context.Context, category db.CreateCategoryParams, course db.CreateCourseParams) error {
	return c.callTx(ctx, func(q *db.Queries) error {
		err := q.CreateCategory(ctx, category)
		if err != nil {
			return err
		}

		course.CategoryID = category.ID
		err = q.CreateCourse(ctx, course)
		if err != nil {
			return err
		}

		return nil
	})
}

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	categoryParams := db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend",
		Description: sql.NullString{String: "Backend Course", Valid: true},
	}

	courseParams := db.CreateCourseParams{
		ID:          uuid.New().String(),
		Name:        "Go",
		Description: "Go Course",
		Price:       10.95,
	}

	courseAR := NewCourseAR(conn)
	err = courseAR.CreateCourseAndCategory(ctx, categoryParams, courseParams)
	if err != nil {
		panic(err)
	}

}
