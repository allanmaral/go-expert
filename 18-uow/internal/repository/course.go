package repository

import (
	"context"
	"database/sql"

	"github.com/allanmaral/go-expert/18-uow/internal/db"
	"github.com/allanmaral/go-expert/18-uow/internal/entity"
)

type CourseRepository interface {
	Insert(ctx context.Context, course entity.Course) error
}

type SQLCourseRepository struct {
	db      *sql.DB
	queries *db.Queries
}

var _ CourseRepository = (*SQLCourseRepository)(nil)

func NewSQLCourseRepository(sdb *sql.DB) *SQLCourseRepository {
	return &SQLCourseRepository{
		db:      sdb,
		queries: db.New(sdb),
	}
}

func (r *SQLCourseRepository) Insert(ctx context.Context, course entity.Course) error {
	return r.queries.CreateCourse(ctx, db.CreateCourseParams{
		Name:       course.Name,
		CategoryID: int32(course.CategoryID),
	})
}
