package service

import (
	"context"
	"io"

	"github.com/allanmaral/go-expert/14-grpc/internal/database"
	"github.com/allanmaral/go-expert/14-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB *database.Category
}

var _ (pb.CategoryServiceServer) = (*CategoryService)(nil)

func NewCategoryService(categoryDB *database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, _ *pb.Void) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	result := make([]*pb.Category, len(categories))
	for i, category := range categories {
		result[i] = &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
	}

	return &pb.CategoryList{Categories: result}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindById(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	result := &pb.CategoryList{}
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(result)
		}
		if err != nil {
			return err
		}

		category, err := c.CategoryDB.Create(in.Name, in.Description)
		if err != nil {
			return err
		}

		result.Categories = append(result.Categories, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
}
