package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *ProductApiServiceServer) GetCategory(ctx context.Context, req *desc.GetCategoryRequest) (*desc.GetCategoryResponse, error) {
	res, err := s.categoryService.GetCategoryByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return s.categoryToProtoGetCategoryResponse(res), nil
}

func (s *ProductApiServiceServer) categoryToProtoGetCategoryResponse(resp *datastruct.Category) *desc.GetCategoryResponse {
	internalResp := &desc.GetCategoryResponse{
		Id:    resp.ID,
		Name:  resp.Name,
		Level: resp.Level,
	}
	if resp.ParentID.Valid {
		internalResp.ParentId = wrapperspb.Int64(resp.ParentID.Int64)
	}
	return internalResp
}
