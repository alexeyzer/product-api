package product_serivce

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alexeyzer/product-api/config"
	"github.com/alexeyzer/product-api/internal/pkg/service"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	userapi "github.com/alexeyzer/product-api/pb/api/user/v1"

	"google.golang.org/grpc/metadata"
)

type ProductApiServiceServer struct {
	brandService        service.BrandService
	categoryService     service.CategoryService
	colorService        service.ColorService
	sizeService         service.SizeService
	productService      service.ProductService
	finalProductService service.FinalProductService
	desc.UnimplementedProductApiServiceServer
}

func (s *ProductApiServiceServer) GetUserInfoFromContext(ctx context.Context) (*userapi.SessionCheckResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	userInfo := &userapi.SessionCheckResponse{}
	if ok {
		val := md.Get(config.Config.Auth.UserInfoKey)
		if len(val) > 0 {
			err := json.Unmarshal([]byte(val[0]), userInfo)
			if err != nil {
				return nil, err
			}
			return userInfo, nil
		}
	}
	return nil, fmt.Errorf("no userInfo in context")
}

func NewProductApiServiceServer(
	categoryService service.CategoryService,
	brandService service.BrandService,
	colorService service.ColorService,
	sizeService service.SizeService,
	productService service.ProductService,
	finalProductService service.FinalProductService,
) *ProductApiServiceServer {
	return &ProductApiServiceServer{
		categoryService:     categoryService,
		brandService:        brandService,
		colorService:        colorService,
		sizeService:         sizeService,
		productService:      productService,
		finalProductService: finalProductService,
	}
}
