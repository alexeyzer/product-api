package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProductApiServiceServer) ListBrandsGrouped(ctx context.Context, _ *emptypb.Empty) (*desc.ListBrandsGroupedResponse, error) {
	resp, err := s.brandService.ListBrandsGrouped(ctx)
	if err != nil {
		return nil, err
	}
	return s.brandGroupToProtoListBrandsGroupedResponse(resp), nil
}

func (s *ProductApiServiceServer) brandGroupToProtoListBrandsGroupedResponse(resp []*datastruct.BrandGroup) *desc.ListBrandsGroupedResponse {
	result := &desc.ListBrandsGroupedResponse{
		BrandGroups: make([]*desc.ListBrandsGroupedResponse_BrandGroup, 0, len(resp)),
	}
	for _, item := range resp {
		group := &desc.ListBrandsGroupedResponse_BrandGroup{
			GroupName: item.GroupName,
			Brands:    make([]*desc.GetBrandResponse, 0, len(item.Brands)),
		}
		for _, brand := range item.Brands {
			group.Brands = append(group.Brands, s.brandToProtoGetBrandResponse(brand))
		}
		result.BrandGroups = append(result.BrandGroups, group)
	}
	return result
}
