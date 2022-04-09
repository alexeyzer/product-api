package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	"github.com/alexeyzer/product-api/internal/pkg/repository"
)

type FinalProductService interface {
	CreateFinalProduct(ctx context.Context, req datastruct.FinalProduct) (*datastruct.FinalProduct, error)
	GetFinalProduct(ctx context.Context, ID int64) (*datastruct.FinalProduct, error)
	DeleteFinalProduct(ctx context.Context, ID int64) error
	ListFinalProducts(ctx context.Context, productID int64) ([]*datastruct.FinalProduct, error)
}

type finalProductService struct {
	dao repository.DAO
}

func (s *finalProductService) DeleteFinalProduct(ctx context.Context, ID int64) error {
	err := s.dao.FinalProductQuery().Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *finalProductService) GetFinalProduct(ctx context.Context, ID int64) (*datastruct.FinalProduct, error) {
	product, err := s.dao.FinalProductQuery().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *finalProductService) ListFinalProducts(ctx context.Context, productID int64) ([]*datastruct.FinalProduct, error) {
	products, err := s.dao.FinalProductQuery().List(ctx, productID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *finalProductService) CreateFinalProduct(ctx context.Context, req datastruct.FinalProduct) (*datastruct.FinalProduct, error) {
	tx, err := s.dao.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = s.dao.FinalProductQuery().Exists(ctx, req.Sku)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "final product with sku = %d already exist", req.Sku)
		}
		return nil, err
	}

	_, err = s.dao.ProductQuery().Get(ctx, req.ProductID)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "product with id = %d doesnt exist", req.ProductID)
		}
		return nil, err
	}

	_, err = s.dao.SizeQuery().Get(ctx, req.SizeID)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.InvalidArgument, "size with id = %d doesnt exist", req.SizeID)
		}
		return nil, err
	}

	finalProduct := datastruct.FinalProduct{
		ProductID: req.ProductID,
		SizeID:    req.SizeID,
		Amount:    req.Amount,
		Sku:       req.Sku,
	}
	createdFinalProduct, err := s.dao.FinalProductQuery().Create(ctx, finalProduct)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return createdFinalProduct, nil
}

func NewFinalProductService(dao repository.DAO) FinalProductService {
	return &finalProductService{dao: dao}
}
