package service

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	"github.com/alexeyzer/product-api/internal/pkg/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SizeService interface {
	CreateSize(ctx context.Context, req datastruct.Size) (*datastruct.Size, error)
	GetSize(ctx context.Context, ID int64) (*datastruct.Size, error)
	DeleteSize(ctx context.Context, ID int64) error
	UpdateSize(ctx context.Context, req datastruct.Size) (*datastruct.Size, error)
	ListSizes(ctx context.Context) ([]*datastruct.SizeWithCategoryName, error)
}

type sizeService struct {
	dao repository.DAO
}

func (s *sizeService) UpdateSize(ctx context.Context, req datastruct.Size) (*datastruct.Size, error) {
	_, err := s.dao.SizeQuery().Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	size, err := s.dao.SizeQuery().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return size, nil
}

func (s *sizeService) DeleteSize(ctx context.Context, ID int64) error {
	err := s.dao.SizeQuery().Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *sizeService) GetSize(ctx context.Context, ID int64) (*datastruct.Size, error) {
	size, err := s.dao.SizeQuery().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return size, nil
}

func (s *sizeService) ListSizes(ctx context.Context) ([]*datastruct.SizeWithCategoryName, error) {
	sizes, err := s.dao.SizeQuery().List(ctx)
	if err != nil {
		return nil, err
	}

	return sizes, nil
}

func (s *sizeService) CreateSize(ctx context.Context, req datastruct.Size) (*datastruct.Size, error) {
	exists, err := s.dao.SizeQuery().Exists(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, status.Errorf(codes.InvalidArgument, "size with name = %s already exist", req.Name)
	}

	_, err = s.dao.CategoryQuery().Get(ctx, req.CategoryID)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "category with id = %d doesn`t exist", req.CategoryID)
		}
		return nil, err
	}

	res, err := s.dao.SizeQuery().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewSizeService(dao repository.DAO) SizeService {
	return &sizeService{dao: dao}
}
