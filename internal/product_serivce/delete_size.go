package product_serivce

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) DeleteSize(ctx context.Context, req *desc.DeleteSizeRequest) (*emptypb.Empty, error) {
	err := s.sizeService.DeleteSize(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
