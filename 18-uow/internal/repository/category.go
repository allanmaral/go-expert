package repository

import (
	"context"
	"database/sql"

	"github.com/allanmaral/go-expert/18-uow/internal/db"
	"github.com/allanmaral/go-expert/18-uow/internal/entity"
)

type CategoryRepository interface {
	Insert(ctx context.Context, course entity.Category) error
}

type SQLCategoryRepository struct {
	db      *sql.DB
	queries *db.Queries
}

var _ CategoryRepository = (*SQLCategoryRepository)(nil)

func NewSQLCategoryRepository(sdb *sql.DB) *SQLCategoryRepository {
	return &SQLCategoryRepository{
		db:      sdb,
		queries: db.New(sdb),
	}
}

func (r *SQLCategoryRepository) Insert(ctx context.Context, course entity.Category) error {
	return r.queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: course.Name,
	})
}
