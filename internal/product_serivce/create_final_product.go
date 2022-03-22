package product_serivce

import (
	"context"

	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) CreateFinalProduct(ctx context.Context, req *desc.CreateFinalProductRequest) (*desc.CreateFinalProductResponse, error) {
	resp, err := s.finalProductService.CreateFinalProduct(ctx, s.protoCreateFinalProductRequestToFinalProduct(req))
	if err != nil {
		return nil, err
	}
	return s.finalProductToProtoCreateFinalProductResponse(resp), nil
}

func (s *ProductApiServiceServer) protoCreateFinalProductRequestToFinalProduct(req *desc.CreateFinalProductRequest) datastruct.FinalProduct {
	return datastruct.FinalProduct{
		ProductID: req.GetProductId(),
		SizeID:    req.GetSizeId(),
		ColorID:   req.GetColorId(),
		Amount:    req.GetAmount(),
		Sku:       req.GetSku(),
		Price:     req.GetPrice(),
	}
}

func (s *ProductApiServiceServer) finalProductToProtoCreateFinalProductResponse(resp *datastruct.FinalProduct) *desc.CreateFinalProductResponse {
	return &desc.CreateFinalProductResponse{
		Id:        resp.ID,
		ProductId: resp.ProductID,
		SizeId:    resp.SizeID,
		ColorId:   resp.ColorID,
		Price:     resp.Price,
		Sku:       resp.Sku,
		Amount:    resp.Amount,
	}
}
