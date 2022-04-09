package product_serivce

import (
	"context"

	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) GetProduct(ctx context.Context, req *desc.GetProductRequest) (*desc.GetProductResponse, error) {
	resp, err := s.productService.GetProduct(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return s.productToProtoGetProductResponse(resp), nil
}

func (s *ProductApiServiceServer) productToProtoGetProductResponse(resp *datastruct.Product) *desc.GetProductResponse {
	return &desc.GetProductResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Description,
		Url:         resp.Url,
		BrandId:     resp.BrandID,
		CategoryId:  resp.CategoryID,
		Color:       resp.Color,
		Price:       resp.Price,
	}
}
