package product_serivce

import (
	"context"
	"database/sql"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) CreateBrand(ctx context.Context, req *desc.CreateBrandRequest) (*desc.CreateBrandResponse, error) {
	res, err := s.brandService.CreateBrand(ctx, s.protoCreateBrandRequestToDatastruct(req), req.GetFile(), req.GetFileExtension())
	if err != nil {
		return nil, err
	}
	return s.datastructToProtoCreateBrandResponse(res), nil
}

func (s *ProductApiServiceServer) protoCreateBrandRequestToDatastruct(req *desc.CreateBrandRequest) datastruct.Brand {
	brand := datastruct.Brand{
		Name: req.GetName(),
	}
	if req.GetDescription() != "" {
		brand.Description = sql.NullString{String: req.GetDescription(), Valid: true}
	}
	return brand
}

func (s *ProductApiServiceServer) datastructToProtoCreateBrandResponse(res *datastruct.Brand) *desc.CreateBrandResponse {
	internalResp := &desc.CreateBrandResponse{
		Id:   res.ID,
		Name: res.Name,
	}
	if res.Url.Valid {
		internalResp.ImageUrl = res.Url.String
	}
	if res.Description.Valid {
		internalResp.Description = res.Description.String
	}
	return internalResp
}
