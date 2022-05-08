package service

import (
	"bytes"
	"context"
	"github.com/alexeyzer/product-api/internal/client"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	"github.com/alexeyzer/product-api/internal/pkg/repository"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req datastruct.CreateProduct) (*datastruct.Product, error)
	GetProduct(ctx context.Context, ID int64) (*datastruct.Product, error)
	GetFullProduct(ctx context.Context, ID int64) (*datastruct.FullProduct, error)
	DeleteProduct(ctx context.Context, ID int64) error
	ListProducts(ctx context.Context, req datastruct.ListProductRequest) ([]*datastruct.Product, error)
	ListProductsByPhoto(ctx context.Context, image []byte) ([]*datastruct.Product, error)
}

type productService struct {
	dao                repository.DAO
	s3                 client.S3Client
	recognizeAPIClient client.RecognizeAPIClient
}

func (s *productService) ListProductsByPhoto(ctx context.Context, image []byte) ([]*datastruct.Product, error) {
	res, err := s.recognizeAPIClient.RecognizePhoto(ctx, image)
	if err != nil {
		return nil, err
	}
	products, err := s.dao.ProductQuery().ListByCategoryID(ctx, res)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productService) GetFullProduct(ctx context.Context, ID int64) (*datastruct.FullProduct, error) {
	product, err := s.dao.ProductQuery().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	sizes, err := s.dao.SizeQuery().GetByProductID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return &datastruct.FullProduct{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Url:         product.Url,
		BrandID:     0,
		CategoryID:  0,
		Price:       product.Price,
		Color:       product.Color,
		Sizes:       sizes,
	}, nil
}

func (s *productService) DeleteProduct(ctx context.Context, ID int64) error {
	tx, err := s.dao.BeginTransaction(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.dao.ProductQuery().Delete(ctx, ID)
	if err != nil {
		return err
	}

	err = s.dao.FinalProductQuery().DeleteByProductID(ctx, ID)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (s *productService) GetProduct(ctx context.Context, ID int64) (*datastruct.Product, error) {
	product, err := s.dao.ProductQuery().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) ListProducts(ctx context.Context, req datastruct.ListProductRequest) ([]*datastruct.Product, error) {
	products, err := s.dao.ProductQuery().List(ctx, req)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productService) CreateProduct(ctx context.Context, req datastruct.CreateProduct) (*datastruct.Product, error) {
	tx, err := s.dao.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = s.dao.ProductQuery().Exists(ctx, req.Name)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "product with name = %s already exist", req.Name)
		}
		return nil, err
	}

	_, err = s.dao.BrandQuery().Get(ctx, req.BrandID)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "brand with ID = % doesn`t exist", req.BrandID)
		}
		return nil, err
	}
	_, err = s.dao.CategoryQuery().Get(ctx, req.CategoryID)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "category with ID = % doesn`t exist", req.CategoryID)
		}
		return nil, err
	}

	product := datastruct.Product{
		Name:        req.Name,
		Description: req.Description,
		BrandID:     req.BrandID,
		CategoryID:  req.CategoryID,
		Price:       req.Price,
		Color:       req.Color,
	}
	if len(req.Image) > 0 && req.ContentType != "" {
		res, err := s.s3.UploadFileD(ctx, uuid.New().String()+req.ContentType, bytes.NewReader(req.Image), req.ContentType)
		if err != nil {
			return nil, err
		}
		product.Url = res.Location
	}
	createdProduct, err := s.dao.ProductQuery().Create(ctx, product)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return createdProduct, nil
}

func NewProductService(dao repository.DAO, s3 client.S3Client, recognizeAPIClient client.RecognizeAPIClient) ProductService {
	return &productService{
		dao:                dao,
		s3:                 s3,
		recognizeAPIClient: recognizeAPIClient,
	}
}
