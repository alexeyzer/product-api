package service

import (
	"context"
	"github.com/alexeyzer/product-api/config"
	"github.com/alexeyzer/product-api/internal/client"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	"github.com/alexeyzer/product-api/internal/pkg/repository"
)

type FinalProductService interface {
	CreateFinalProduct(ctx context.Context, req datastruct.FinalProduct) (*datastruct.FinalProduct, error)
	BatchUpdateFinalProduct(ctx context.Context, req []datastruct.FinalProduct) error
	UpdateFinalProduct(ctx context.Context, req datastruct.FinalProduct) (*datastruct.FinalProduct, error)
	GetFinalProduct(ctx context.Context, ID int64) (*datastruct.FinalProductWithSizeName, error)
	DeleteFinalProduct(ctx context.Context, ID int64) error
	ListFinalProducts(ctx context.Context, productID int64, session string) ([]*datastruct.FinalProductWithSizeName, error)
	ListFullFinalProducts(ctx context.Context, finalProductIds []int64) ([]*datastruct.FullFinalProduct, error)
}

type finalProductService struct {
	dao           repository.DAO
	userAPIClient client.UserAPIClient
}

func (s *finalProductService) BatchUpdateFinalProduct(ctx context.Context, req []datastruct.FinalProduct) error {
	err := s.dao.FinalProductQuery().ButchUpdate(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (s *finalProductService) UpdateFinalProduct(ctx context.Context, req datastruct.FinalProduct) (*datastruct.FinalProduct, error) {
	_, err := s.dao.FinalProductQuery().Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	resp, err := s.dao.FinalProductQuery().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *finalProductService) DeleteFinalProduct(ctx context.Context, ID int64) error {
	err := s.dao.FinalProductQuery().Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *finalProductService) GetFinalProduct(ctx context.Context, ID int64) (*datastruct.FinalProductWithSizeName, error) {
	product, err := s.dao.FinalProductQuery().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *finalProductService) ListFinalProducts(ctx context.Context, productID int64, session string) ([]*datastruct.FinalProductWithSizeName, error) {
	products, err := s.dao.FinalProductQuery().List(ctx, productID)
	if err != nil {
		return nil, err
	}
	if session != "" {
		ctx = metadata.AppendToOutgoingContext(ctx, config.Config.Auth.SessionKey, session)
		resp, err := s.userAPIClient.ListCartItems(ctx)
		if err != nil {
			return nil, err
		}
		quantityMap := make(map[int64]int64)
		for _, item := range resp.GetProducts() {
			quantityMap[item.FullProductId] = item.UserQuantity
		}
		for _, product := range products {
			if quantity, ok := quantityMap[product.ID]; ok {
				product.UserQuantity = quantity
			}
		}
	}
	return products, nil
}

func (s *finalProductService) ListFullFinalProducts(ctx context.Context, finalProductIds []int64) ([]*datastruct.FullFinalProduct, error) {
	products, err := s.dao.FinalProductQuery().ListFull(ctx, finalProductIds)
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

func NewFinalProductService(dao repository.DAO, userAPIClient client.UserAPIClient) FinalProductService {
	return &finalProductService{
		dao:           dao,
		userAPIClient: userAPIClient,
	}
}
