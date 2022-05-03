package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) ListFullFinalProducts(ctx context.Context, req *desc.ListFullFinalProductsRequest) (*desc.ListFullFinalProductsResponse, error) {
	resp, err := s.finalProductService.ListFullFinalProducts(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}
	return s.datastructFullFInalProductsToProto(resp), nil
}

func (s *ProductApiServiceServer) datastructFullFInalProductsToProto(resp []*datastruct.FullFinalProduct) *desc.ListFullFinalProductsResponse {
	internalResp := &desc.ListFullFinalProductsResponse{
		Products: make([]*desc.ListFullFinalProductsResponse_FullFinalProduct, 0, len(resp)),
	}
	for _, product := range resp {
		internalResp.Products = append(internalResp.Products, &desc.ListFullFinalProductsResponse_FullFinalProduct{
			Id:           product.ID,
			Name:         product.Name,
			Description:  product.Description,
			Url:          product.Url,
			BrandName:    product.BrandName,
			CategoryName: product.CategoryName,
			Color:        product.Color,
			Price:        product.Price,
			Size:         product.Size,
			Amount:       product.Amount,
			Sku:          product.Sku,
		})
	}
	return internalResp
}
