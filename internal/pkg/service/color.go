package service

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	"github.com/alexeyzer/product-api/internal/pkg/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ColorService interface {
	CreateColor(ctx context.Context, req datastruct.Color) (*datastruct.Color, error)
	GetColor(ctx context.Context, ID int64) (*datastruct.Color, error)
	DeleteColor(ctx context.Context, ID int64) error
	ListColors(ctx context.Context) ([]*datastruct.Color, error)
}

type colorService struct {
	dao repository.DAO
}

func (s *colorService) DeleteColor(ctx context.Context, ID int64) error {
	err := s.dao.ColorQuery().Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *colorService) GetColor(ctx context.Context, ID int64) (*datastruct.Color, error) {
	color, err := s.dao.ColorQuery().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return color, nil
}

func (s *colorService) ListColors(ctx context.Context) ([]*datastruct.Color, error) {
	colors, err := s.dao.ColorQuery().List(ctx)
	if err != nil {
		return nil, err
	}

	return colors, nil
}

func (s *colorService) CreateColor(ctx context.Context, req datastruct.Color) (*datastruct.Color, error) {
	_, err := s.dao.ColorQuery().Exists(ctx, req.Name)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "color with name = %s already exist", req.Name)
		}
		return nil, err
	}

	res, err := s.dao.ColorQuery().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewColorService(dao repository.DAO) ColorService {
	return &colorService{dao: dao}
}
