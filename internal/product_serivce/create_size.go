package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) CreateSize(ctx context.Context, req *desc.CreateSizeRequest) (*desc.CreateSizeResponse, error) {
	res, err := s.sizeService.CreateSize(ctx, s.protoCreateSizeRequestToSize(req))
	if err != nil {
		return nil, err
	}
	return s.sizeToProtoCreateSizeResponse(res), nil
}

func (s *ProductApiServiceServer) protoCreateSizeRequestToSize(req *desc.CreateSizeRequest) datastruct.Size {
	return datastruct.Size{
		Name:       req.GetName(),
		CategoryID: req.GetCategoryId(),
	}
}

func (s *ProductApiServiceServer) sizeToProtoCreateSizeResponse(resp *datastruct.Size) *desc.CreateSizeResponse {
	return &desc.CreateSizeResponse{
		Id:         resp.ID,
		Name:       resp.Name,
		CategoryId: resp.CategoryID,
	}
}
