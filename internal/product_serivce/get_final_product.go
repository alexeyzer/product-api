package product_serivce

import (
	"context"

	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) GetFinalProduct(ctx context.Context, req *desc.GetFinalProductRequest) (*desc.GetFinalProductResponse, error) {
	resp, err := s.finalProductService.GetFinalProduct(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return s.finalProductToProtoGetFinalProductResponse(resp), nil
}

func (s *ProductApiServiceServer) finalProductToProtoGetFinalProductResponse(resp *datastruct.FinalProduct) *desc.GetFinalProductResponse {
	return &desc.GetFinalProductResponse{
		Id:        resp.ID,
		ProductId: resp.ProductID,
		SizeId:    resp.SizeID,
		Sku:       resp.Sku,
		Amount:    resp.Amount,
	}
}
