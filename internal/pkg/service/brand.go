package service

import (
	"bytes"
	"context"
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexeyzer/product-api/internal/client"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	"github.com/alexeyzer/product-api/internal/pkg/repository"
)

type BrandService interface {
	CreateBrand(ctx context.Context, req datastruct.Brand, image []byte, contentType string) (*datastruct.Brand, error)
	GetBrand(ctx context.Context, ID int64) (*datastruct.Brand, error)
	DeleteBrand(ctx context.Context, ID int64) error
	ListBrands(ctx context.Context) ([]*datastruct.BrandGroup, error)
}

type brandService struct {
	dao repository.DAO
	s3  client.S3Client
}

func (s *brandService) DeleteBrand(ctx context.Context, ID int64) error {
	err := s.dao.BrandQuery().Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *brandService) GetBrand(ctx context.Context, ID int64) (*datastruct.Brand, error) {
	brand, err := s.dao.BrandQuery().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return brand, nil
}

func (s *brandService) ListBrands(ctx context.Context) ([]*datastruct.BrandGroup, error) {
	brands, err := s.dao.BrandQuery().List(ctx)
	if err != nil {
		return nil, err
	}

	brandGroups := make([]*datastruct.BrandGroup, 0)
	if len(brands) == 0 {
		return brandGroups, nil
	}
	prevItemGroup := string(brands[0].Name[0])
	brandGroup := &datastruct.BrandGroup{
		GroupName: prevItemGroup,
		Brands:    make([]*datastruct.Brand, 0),
	}
	for _, item := range brands {
		if strings.ToUpper(string(item.Name[0])) != prevItemGroup {
			brandGroups = append(brandGroups, brandGroup)
			prevItemGroup = string(item.Name[0])
			brandGroup = &datastruct.BrandGroup{
				GroupName: string(item.Name[0]),
				Brands:    make([]*datastruct.Brand, 0),
			}
			brandGroup.Brands = append(brandGroup.Brands, item)
		} else {
			brandGroup.Brands = append(brandGroup.Brands, item)
		}
	}
	brandGroups = append(brandGroups, brandGroup)
	return brandGroups, nil
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
