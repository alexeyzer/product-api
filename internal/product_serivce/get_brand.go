package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) GetBrand(ctx context.Context, req *desc.GetBrandRequest) (*desc.GetBrandResponse, error) {
	resp, err := s.brandService.GetBrand(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return s.brandToProtoGetBrandResponse(resp), nil
}

func (s *ProductApiServiceServer) brandToProtoGetBrandResponse(resp *datastruct.Brand) *desc.GetBrandResponse {
	internalResp := &desc.GetBrandResponse{
		Id:          resp.ID,
		Name:        resp.Name,
		Description: resp.Description.String,
		ImageUrl:    resp.Url.String,
	}
	return internalResp
}
