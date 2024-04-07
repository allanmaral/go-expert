package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/allanmaral/go-expert/18-uow/internal/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestAddCourse(t *testing.T) {
	// This is bad
	sdb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	assert.NoError(t, err)

	// This is really bad!
	sdb.Exec("DROP TABLE if exists `courses`;")
	sdb.Exec("DROP TABLE if exists `categories`;")

	sdb.Exec("CREATE TABLE IF NOT EXISTS `categories` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL);")
	sdb.Exec("CREATE TABLE IF NOT EXISTS `courses` (id int PRIMARY KEY AUTO_INCREMENT, name varchar(255) NOT NULL, category_id INTEGER NOT NULL, FOREIGN KEY (category_id) REFERENCES categories(id));")

	input := AddCourseInput{
		CategoryName:     "Category 1",
		CourseName:       "Course 1",
		CourseCategoryID: 1,
	}

	ctx := context.Background()

	uc := NewAddCourseUseCase(repository.NewSQLCourseRepository(sdb), repository.NewSQLCategoryRepository(sdb))
	err = uc.Execute(ctx, input)
	assert.NoError(t, err)
}
