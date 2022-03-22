package product_serivce

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) DeleteProduct(ctx context.Context, req *desc.DeleteProductRequest) (*emptypb.Empty, error) {
	err := s.productService.DeleteProduct(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
