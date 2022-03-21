package product_serivce

import (
	"context"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProductApiServiceServer) DeleteColor(ctx context.Context, req *desc.DeleteColorRequest) (*emptypb.Empty, error) {
	err := s.colorService.DeleteColor(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
