package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (s *ProductApiServiceServer) ListCategory(ctx context.Context, req *desc.ListCategoryRequest) (*desc.ListCategoryResponse, error) {
	res, err := s.categoryService.ListCategory(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.categoriesToProtoListCategoryResponse(res), nil
}

func (s *ProductApiServiceServer) categoriesToProtoListCategoryResponse(resp []datastruct.Category) *desc.ListCategoryResponse {
	internalResp := &desc.ListCategoryResponse{
		Items: make([]*desc.Category, 0, len(resp)),
	}
	for _, category := range resp {
		internalCategory := &desc.Category{
			Id:    category.ID,
			Name:  category.Name,
			Level: category.Level,
		}
		if category.ParentID.Valid {
			internalCategory.ParentId = wrapperspb.Int64(category.ParentID.Int64)
		}
		internalResp.Items = append(internalResp.Items, internalCategory)
	}
	return internalResp
}
