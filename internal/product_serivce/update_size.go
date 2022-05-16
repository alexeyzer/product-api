package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"

	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) UpdateSize(ctx context.Context, req *desc.UpdateSizeRequest) (*desc.UpdateSizeResponse, error) {
	res, err := s.sizeService.UpdateSize(ctx, datastruct.Size{
		ID:         req.GetId(),
		Name:       req.GetName(),
		CategoryID: req.GetCategoryId(),
	})
	if err != nil {
		return nil, err
	}
	return &desc.UpdateSizeResponse{
		Id:         res.ID,
		Name:       res.Name,
		CategoryId: res.CategoryID,
	}, nil
}
