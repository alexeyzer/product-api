package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProductApiServiceServer) ListSizes(ctx context.Context, _ *emptypb.Empty) (*desc.ListSizesResponse, error) {
	resp, err := s.sizeService.ListSizes(ctx)
	if err != nil {
		return nil, err
	}
	return s.sizesToProtoListSizesResponse(resp), nil
}

func (s *ProductApiServiceServer) sizesToProtoListSizesResponse(resp []*datastruct.Size) *desc.ListSizesResponse {
	internalResp := &desc.ListSizesResponse{
		Sizes: make([]*desc.CreateSizeResponse, 0, len(resp)),
	}
	for _, item := range resp {
		internalResp.Sizes = append(internalResp.Sizes, &desc.CreateSizeResponse{
			Id:       item.ID,
			Name:     item.Name,
			Category: item.Category,
		})
	}
	return internalResp
}
