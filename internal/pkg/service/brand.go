package service

import (
	"bytes"
	"context"
	"database/sql"
	"github.com/alexeyzer/product-api/internal/client"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	"github.com/alexeyzer/product-api/internal/pkg/repository"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BrandService interface {
	CreateBrand(ctx context.Context, req datastruct.Brand, image []byte, contentType string) (*datastruct.Brand, error)
}

type brandService struct {
	dao repository.DAO
	s3  client.S3Client
}

func (s *brandService) CreateBrand(ctx context.Context, req datastruct.Brand, image []byte, contentType string) (*datastruct.Brand, error) {
	_, err := s.dao.BrandQuery().Exists(ctx, req.Name)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "brand with name = %s already exist", req.Name)
		}
		return nil, err
	}
	if len(image) > 0 && contentType != "" {
		res, err := s.s3.UploadFileD(ctx, uuid.New().String(), bytes.NewReader(image), contentType)
		if err != nil {
			return nil, err
		}
		req.Url = sql.NullString{String: res.Location, Valid: true}
	}
	res, err := s.dao.BrandQuery().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewBrandService(dao repository.DAO, s3 client.S3Client) BrandService {
	return &brandService{dao: dao, s3: s3}
}
