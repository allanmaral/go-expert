package usecase

import (
	"context"

	"github.com/allanmaral/go-expert/18-uow/internal/entity"
	"github.com/allanmaral/go-expert/18-uow/internal/repository"
)

type AddCourseInput struct {
	CategoryName     string
	CourseName       string
	CourseCategoryID int
}

type AddCourseUseCase struct {
	courseRepository   repository.CourseRepository
	categoryRepository repository.CategoryRepository
}

func NewAddCourseUseCase(courseRepository repository.CourseRepository, categoryRepository repository.CategoryRepository) *AddCourseUseCase {
	return &AddCourseUseCase{
		courseRepository:   courseRepository,
		categoryRepository: categoryRepository,
	}
}

func (uc *AddCourseUseCase) Execute(ctx context.Context, input AddCourseInput) error {
	category := entity.Category{
		Name: input.CategoryName,
	}

	err := uc.categoryRepository.Insert(ctx, category)
	if err != nil {
		return err
	}

	course := entity.Course{
		Name:       input.CourseName,
		CategoryID: input.CourseCategoryID,
	}

	err = uc.courseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}

	return nil
}
