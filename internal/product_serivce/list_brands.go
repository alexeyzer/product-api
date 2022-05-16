package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) ListBrands(ctx context.Context, _ *emptypb.Empty) (*desc.ListBrandsResponse, error) {
	resp, err := s.brandService.ListBrands(ctx)
	if err != nil {
		return nil, err
	}
	return s.brandsToProtoListBrandsResponse(resp), nil
}

func (s *ProductApiServiceServer) brandsToProtoListBrandsResponse(brands []*datastruct.Brand) *desc.ListBrandsResponse {
	resp := &desc.ListBrandsResponse{
		Brands: make([]*desc.GetBrandResponse, 0, len(brands)),
	}
	for _, brand := range brands {
		resp.Brands = append(resp.Brands, s.brandToProtoGetBrandResponse(brand))
	}
	return resp
}
