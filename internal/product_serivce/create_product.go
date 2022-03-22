package product_serivce

import (
	"context"

	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) CreateProduct(ctx context.Context, req *desc.CreateProductRequest) (*desc.CreateProductResponse, error) {
	resp, err := s.productService.CreateProduct(ctx, s.protoCreateProductRequestToProduct(req))
	if err != nil {
		return nil, err
	}
	return s.productToProtoCreateProductResponse(resp), nil
}

func (s *ProductApiServiceServer) protoCreateProductRequestToProduct(req *desc.CreateProductRequest) datastruct.CreateProduct {
	return datastruct.CreateProduct{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Image:       req.GetImage(),
		ContentType: req.GetContentType(),
		BrandID:     req.GetBrandId(),
		CategoryID:  req.GetCategoryId(),
	}
}

func (s *ProductApiServiceServer) productToProtoCreateProductResponse(resp *datastruct.Product) *desc.CreateProductResponse {
	return &desc.CreateProductResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Description,
		BrandId:     resp.BrandID,
		CategoryId:  resp.CategoryID,
		Url:         resp.Url,
	}
}
