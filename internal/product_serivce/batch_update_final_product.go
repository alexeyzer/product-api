package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProductApiServiceServer) BatchUpdateFinalProduct(ctx context.Context, req *desc.BatchUpdateFinalProductRequest) (*emptypb.Empty, error) {
	internalReq := make([]datastruct.FinalProduct, 0, len(req.GetItems()))
	for _, item := range req.GetItems() {
		internalReq = append(internalReq, datastruct.FinalProduct{
			ID:     item.GetId(),
			Amount: item.GetAmount(),
		})
	}

	err := s.finalProductService.BatchUpdateFinalProduct(ctx, internalReq)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
