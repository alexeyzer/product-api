package service

import (
	"context"
	"database/sql"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	"github.com/alexeyzer/product-api/internal/pkg/repository"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, req *desc.CreateCategoryRequest) (*datastruct.Category, error)
	GetCategoryByID(ctx context.Context, ID int64) (*datastruct.Category, error)
	DeleteCategory(ctx context.Context, ID int64) error
	UpdateCategory(ctx context.Context, req *desc.UpdateCategoryRequest) (*datastruct.Category, error)
	ListCategory(ctx context.Context, req *desc.ListCategoryRequest) ([]datastruct.Category, error)
}

type categoryService struct {
	dao repository.DAO
}

func (s *categoryService) ListCategory(ctx context.Context, req *desc.ListCategoryRequest) ([]datastruct.Category, error) {
	res, err := s.dao.CategoryQuery().List(ctx, s.protoListCategoryToDatastruct(req))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *categoryService) protoListCategoryToDatastruct(req *desc.ListCategoryRequest) datastruct.ListCategoryRequest {
	internalReq := datastruct.ListCategoryRequest{
		Offset: req.Page.Limit * (req.Page.Number - 1),
		Limit:  req.Page.Limit,
	}
	if req.Name != nil {
		internalReq.Name = sql.NullString{String: req.Name.Value, Valid: true}
	}
	if req.Level != nil {
		internalReq.Level = sql.NullInt64{Int64: req.Level.Value, Valid: true}
	}
	return internalReq
}

func (s *categoryService) GetCategoryByID(ctx context.Context, ID int64) (*datastruct.Category, error) {
	category, err := s.dao.CategoryQuery().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) CreateCategory(ctx context.Context, req *desc.CreateCategoryRequest) (*datastruct.Category, error) {
	exists, err := s.dao.CategoryQuery().Exists(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	if exists == true {
		return nil, status.Errorf(codes.InvalidArgument, "Category with name = %s already exists", req.Name)
	}
	res, err := s.dao.CategoryQuery().Create(ctx, s.serviceCreateCategoryReqToDaoCategory(req))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *categoryService) UpdateCategory(ctx context.Context, req *desc.UpdateCategoryRequest) (*datastruct.Category, error) {
	_, err := s.dao.CategoryQuery().Get(ctx, req.GetId())
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "Category with id = %d doesn't exist", req.GetId())
		}
		return nil, err
	}
	res, err := s.dao.CategoryQuery().Update(ctx, s.serviceUpdateReqToDaoCategory(req))
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *categoryService) serviceUpdateReqToDaoCategory(req *desc.UpdateCategoryRequest) datastruct.Category {
	internalCategory := datastruct.Category{
		Name:  req.Name,
		Level: req.Level,
	}
	if req.ParentId != nil {
		internalCategory.ParentID = sql.NullInt64{Int64: req.ParentId.Value, Valid: true}
	}
	return internalCategory
}

func (s *categoryService) DeleteCategory(ctx context.Context, ID int64) error {
	_, err := s.dao.CategoryQuery().Get(ctx, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.InvalidArgument, "Category with id = %d doesn't exist", ID)
		}
		return err
	}
	err = s.dao.CategoryQuery().Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *categoryService) serviceCreateCategoryReqToDaoCategory(req *desc.CreateCategoryRequest) datastruct.Category {
	internalCategory := datastruct.Category{
		Name:  strings.ToLower(req.Name),
		Level: req.Level,
	}
	if req.ParentId != nil {
		internalCategory.ParentID = sql.NullInt64{Int64: req.ParentId.Value, Valid: true}
	}
	return internalCategory
}

func NewCategoryService(dao repository.DAO) CategoryService {
	return &categoryService{dao: dao}
}
