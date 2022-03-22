package product_serivce

import (
	"context"

	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) GetSize(ctx context.Context, req *desc.GetSizeRequest) (*desc.GetSizeResponse, error) {
	res, err := s.sizeService.GetSize(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return s.sizeToProtoGetSizeResponse(res), nil
}

func (s *ProductApiServiceServer) sizeToProtoGetSizeResponse(resp *datastruct.Size) *desc.GetSizeResponse {
	return &desc.GetSizeResponse{
		Id:         resp.ID,
		Name:       resp.Name,
		CategoryId: resp.CategoryID,
	}
}
