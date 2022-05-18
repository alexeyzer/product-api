package product_serivce

import (
	"context"
	"database/sql"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) ListProducts(ctx context.Context, req *desc.ListProductsRequest) (*desc.ListProductsResponse, error) {
	resp, err := s.productService.ListProducts(ctx, s.listProductsRequestToListRequest(req))
	if err != nil {
		return nil, err
	}
	return s.productsToProtoListProductsResponse(resp), nil
}

func (s *ProductApiServiceServer) listProductsRequestToListRequest(req *desc.ListProductsRequest) datastruct.ListProductRequest {
	internalReq := datastruct.ListProductRequest{
		Offset: req.GetPage().GetLimit() * (req.GetPage().GetNumber() - 1),
		Limit:  req.GetPage().GetLimit(),
		IsAll:  req.GetPage().GetIsAll(),
		Name:   req.GetName(),
	}
	if req.BrandId != nil {
		internalReq.BrandID = sql.NullInt64{Int64: req.BrandId.Value, Valid: true}
	}
	if req.CategoryId != nil {
		internalReq.CategoryID = sql.NullInt64{Int64: req.CategoryId.Value, Valid: true}
	}
	return internalReq
}

func (s *ProductApiServiceServer) productsToProtoListProductsResponse(resp []*datastruct.Product) *desc.ListProductsResponse {
	internalResp := &desc.ListProductsResponse{
		Products: make([]*desc.GetProductResponse, 0, len(resp)),
	}
	for _, item := range resp {
		internalResp.Products = append(internalResp.Products, s.productToProtoGetProductResponse(item))
	}
	return internalResp
}
