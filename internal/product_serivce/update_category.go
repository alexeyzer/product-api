package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *ProductApiServiceServer) UpdateCategory(ctx context.Context, req *desc.UpdateCategoryRequest) (*desc.UpdateCategoryResponse, error) {
	res, err := s.categoryService.UpdateCategory(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.categoryToProtoUpdateCategoryResponse(res), nil
}

func (s *ProductApiServiceServer) categoryToProtoUpdateCategoryResponse(resp *datastruct.Category) *desc.UpdateCategoryResponse {
	internalResp := &desc.UpdateCategoryResponse{
		Id:    resp.ID,
		Name:  resp.Name,
		Level: resp.Level,
	}
	if resp.ParentID.Valid {
		internalResp.ParentId = wrapperspb.Int64(resp.ParentID.Int64)
	}
	return internalResp
}
