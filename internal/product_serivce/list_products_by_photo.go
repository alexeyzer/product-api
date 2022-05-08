package product_serivce

import (
	"context"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	log "github.com/sirupsen/logrus"
)

func (s *ProductApiServiceServer) ListProductsByPhoto(ctx context.Context, req *desc.ListProductsByPhotoRequest) (*desc.ListProductsResponse, error) {
	log.Printf("byte: %v", req.GetImage())
	res, err := s.productService.ListProductsByPhoto(ctx, req.GetImage())
	if err != nil {
		return nil, err
	}
	return s.productsToProtoListProductsResponse(res), nil
}
