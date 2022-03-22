package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) GetFullProduct(ctx context.Context, req *desc.GetFullProductRequest) (*desc.GetFullProductResponse, error) {
	resp, err := s.productService.GetFullProduct(ctx, req.GetProductId())
	if err != nil {
		return nil, err
	}
	return s.fullProductToProtoGetFullProductResponse(resp), nil
}

func (s *ProductApiServiceServer) fullProductToProtoGetFullProductResponse(resp *datastruct.FullProduct) *desc.GetFullProductResponse {
	internalResp := &desc.GetFullProductResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Description,
		Url:         resp.Url,
		BrandId:     resp.BrandID,
		CategoryId:  resp.CategoryID,
		Sizes:       make([]*desc.GetSizeResponse, 0, len(resp.Sizes)),
		Colors:      make([]*desc.GetColorResponse, 0, len(resp.Colors)),
	}
	for _, item := range resp.Colors {
		internalResp.Colors = append(internalResp.Colors, &desc.GetColorResponse{
			Id:   item.ID,
			Name: item.Name,
		})
	}
	for _, item := range resp.Sizes {
		internalResp.Sizes = append(internalResp.Sizes, &desc.GetSizeResponse{
			Id:         item.ID,
			Name:       item.Name,
			CategoryId: item.CategoryID,
		})
	}
	return internalResp
}
