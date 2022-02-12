package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *ProductApiServiceServer) CreateCategory(ctx context.Context, req *desc.CreateCategoryRequest) (*desc.CreateCategoryResponse, error) {
	res, err := s.categoryService.CreateCategory(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.categoryToProtoCreateCategoryResponse(res), nil
}

func (s *ProductApiServiceServer) categoryToProtoCreateCategoryResponse(resp *datastruct.Category) *desc.CreateCategoryResponse {
	internalResp := &desc.CreateCategoryResponse{
		Id:    resp.ID,
		Name:  resp.Name,
		Level: resp.Level,
	}
	if resp.ParentID.Valid {
		internalResp.ParentId = wrapperspb.Int64(resp.ParentID.Int64)
	}
	return internalResp
}
