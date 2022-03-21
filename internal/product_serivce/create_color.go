package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) CreateColor(ctx context.Context, req *desc.CreateColorRequest) (*desc.CreateColorResponse, error) {
	res, err := s.colorService.CreateColor(ctx, s.protoCreateColorRequestToColor(req))
	if err != nil {
		return nil, err
	}
	return s.colorToProtoCreateColorResponse(res), nil
}

func (s *ProductApiServiceServer) protoCreateColorRequestToColor(req *desc.CreateColorRequest) datastruct.Color {
	return datastruct.Color{
		Name: req.GetName(),
	}
}

func (s *ProductApiServiceServer) colorToProtoCreateColorResponse(resp *datastruct.Color) *desc.CreateColorResponse {
	return &desc.CreateColorResponse{
		Id:   resp.ID,
		Name: resp.Name,
	}
}
