package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) UpdateFinalProduct(ctx context.Context, req *desc.UpdateFinalProductRequest) (*desc.UpdateFinalProductResponse, error) {
	resp, err := s.finalProductService.UpdateFinalProduct(ctx, datastruct.FinalProduct{
		ID:     req.GetId(),
		SizeID: req.GetSizeId(),
		Amount: req.GetAmount(),
		Sku:    req.GetSku(),
	})
	if err != nil {
		return nil, err
	}
	return &desc.UpdateFinalProductResponse{
		Id:        resp.ID,
		ProductId: resp.ProductID,
		SizeId:    resp.SizeID,
		Sku:       resp.Sku,
		Amount:    resp.Amount,
	}, nil
}
