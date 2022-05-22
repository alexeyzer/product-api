package product_serivce

import (
	"context"

	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) ListProductsById(ctx context.Context, req *desc.ListProductsByIdRequest) (*desc.ListProductsResponse, error) {
	resp, err := s.productService.ListProductsByIds(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return s.productsToProtoListProductsResponse(resp), nil
}
