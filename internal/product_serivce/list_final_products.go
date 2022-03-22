package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) ListFinalProducts(ctx context.Context, req *desc.ListFinalProductsRequest) (*desc.ListFinalProductsResponse, error) {
	resp, err := s.finalProductService.ListFinalProducts(ctx, req.GetProductId())
	if err != nil {
		return nil, err
	}
	return s.productsToProtoListFinalProductsResponse(resp), nil
}

func (s *ProductApiServiceServer) productsToProtoListFinalProductsResponse(resp []*datastruct.FinalProduct) *desc.ListFinalProductsResponse {
	internalResp := &desc.ListFinalProductsResponse{
		Products: make([]*desc.GetFinalProductResponse, 0, len(resp)),
	}
	for _, item := range resp {
		internalResp.Products = append(internalResp.Products, s.finalProductToProtoGetFinalProductResponse(item))
	}
	return internalResp
}
