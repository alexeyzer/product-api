package product_serivce

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
)

func (s *ProductApiServiceServer) ListColors(ctx context.Context, _ *emptypb.Empty) (*desc.ListColorsResponse, error) {
	colors, err := s.colorService.ListColors(ctx)
	if err != nil {
		return nil, err
	}
	return s.colorsToProtoListColorsResponse(colors), nil
}

func (s *ProductApiServiceServer) colorsToProtoListColorsResponse(resp []*datastruct.Color) *desc.ListColorsResponse {
	internalResp := &desc.ListColorsResponse{
		Colors: make([]*desc.CreateColorResponse, 0, len(resp)),
	}
	for _, item := range resp {
		internalResp.Colors = append(internalResp.Colors, &desc.CreateColorResponse{
			Id:   item.ID,
			Name: item.Name,
		})
	}
	return internalResp
}
