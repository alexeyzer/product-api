package product_serivce

import (
	"context"

	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) GetColor(ctx context.Context, req *desc.GetColorRequest) (*desc.GetColorResponse, error) {
	resp, err := s.colorService.GetColor(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return s.colorToProtoGetColorResponse(resp), nil
}

func (s *ProductApiServiceServer) colorToProtoGetColorResponse(resp *datastruct.Color) *desc.GetColorResponse {
	return &desc.GetColorResponse{
		Id:   resp.ID,
		Name: resp.Name,
	}
}
