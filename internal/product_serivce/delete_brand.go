package product_serivce

import (
	"context"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProductApiServiceServer) DeleteBrand(ctx context.Context, req *desc.DeleteBrandRequest) (*emptypb.Empty, error) {
	err := s.brandService.DeleteBrand(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
