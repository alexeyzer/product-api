package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ProductApiServiceServer) ListBrands(ctx context.Context, _ *emptypb.Empty) (*desc.ListBrandsResponse, error) {
	resp, err := s.brandService.ListBrands(ctx)
	if err != nil {
		return nil, err
	}
	return s.brandGroupToProtoListBrandsResponse(resp), nil
}
func (s *ProductApiServiceServer) brandGroupToProtoListBrandsResponse(resp []*datastruct.BrandGroup) *desc.ListBrandsResponse {
	result := &desc.ListBrandsResponse{
		BrandGroups: make([]*desc.ListBrandsResponse_BrandGroup, 0, len(resp)),
	}
	for _, item := range resp {
		group := &desc.ListBrandsResponse_BrandGroup{
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
