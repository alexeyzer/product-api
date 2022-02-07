package product_serivce

import (
	"context"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) Echo(ctx context.Context, req *desc.EchoRequest) (*desc.EchoResponse, error) {
	return &desc.EchoResponse{Message: req.Message + " works!"}, nil
}
