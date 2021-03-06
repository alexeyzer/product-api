package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) ListFinalProducts(ctx context.Context, req *desc.ListFinalProductsRequest) (*desc.ListFinalProductsResponse, error) {
	session := s.GetSessionIDFromContext(ctx)
	resp, err := s.finalProductService.ListFinalProducts(ctx, req.GetProductId(), session)
	if err != nil {
		return nil, err
	}
	return s.productsToProtoListFinalProductsResponse(resp), nil
}

func (s *ProductApiServiceServer) productsToProtoListFinalProductsResponse(resp []*datastruct.FinalProductWithSizeName) *desc.ListFinalProductsResponse {
	internalResp := &desc.ListFinalProductsResponse{
		Products: make([]*desc.ListFinalProductsResponse_Item, 0, len(resp)),
	}
	for _, item := range resp {
		internalResp.Products = append(internalResp.Products, &desc.ListFinalProductsResponse_Item{
			Id:           item.ID,
			ProductId:    item.ProductID,
			SizeId:       item.SizeID,
			Sku:          item.Sku,
			Amount:       item.Amount,
			SizeName:     item.SizeName,
			UserQuantity: item.UserQuantity,
		})
	}
	return internalResp
}
