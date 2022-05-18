package product_serivce

import (
	"context"
	"database/sql"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) UpdateBrand(ctx context.Context, req *desc.UpdateBrandRequest) (*desc.UpdateBrandResponse, error) {
	internalReq := datastruct.UpdateBrand{
		ID: req.GetId(),
	}
	if req.GetName() != nil {
		internalReq.Name = sql.NullString{String: req.GetName().GetValue(), Valid: true}
	}
	if req.GetDescription() != nil {
		internalReq.Description = sql.NullString{String: req.GetDescription().GetValue(), Valid: true}
	}
	res, err := s.brandService.UpdateBrand(ctx, internalReq, req.GetFile().GetValue(), req.GetFileExtension().GetValue(), req.GetDeletePhoto())
	if err != nil {
		return nil, err
	}

	resp := &desc.UpdateBrandResponse{
		Id:          res.ID,
		Name:        res.Name,
		Description: res.Description.String,
		ImageUrl:    res.Url.String,
	}
	return resp, nil
}
