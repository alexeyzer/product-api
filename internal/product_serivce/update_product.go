package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"

	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) UpdateProduct(ctx context.Context, req *desc.UpdateProductRequest) (*desc.UpdateProductResponse, error) {
	resp, err := s.productService.UpdateProduct(ctx, datastruct.UpdateProduct{
		ID:          req.GetId(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		BrandID:     req.GetBrandId(),
		CategoryID:  req.GetCategoryId(),
		Price:       req.GetPrice(),
		Color:       req.GetColor(),
	}, req.GetImage(), req.GetContentType(), req.GetDeletePhoto())
	if err != nil {
		return nil, err
	}
	return &desc.UpdateProductResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Description,
		Url:         resp.Url,
		BrandId:     resp.BrandID,
		CategoryId:  resp.CategoryID,
		Color:       resp.Color,
		Price:       resp.Price,
	}, nil
}
