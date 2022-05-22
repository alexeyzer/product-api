package product_serivce

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alexeyzer/product-api/config"
	"github.com/alexeyzer/product-api/internal/pkg/service"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"
	userapi "github.com/alexeyzer/product-api/pb/api/user/v1"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc/metadata"
)

type ProductApiServiceServer struct {
	brandService        service.BrandService
	categoryService     service.CategoryService
	sizeService         service.SizeService
	productService      service.ProductService
	finalProductService service.FinalProductService
	desc.UnimplementedProductApiServiceServer
}

func (s *ProductApiServiceServer) GetSessionIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		val := md.Get(config.Config.Auth.SessionKey)
		if len(val) > 0 {
			return val[0]
		}
		log.Info("no value with key:", config.Config.Auth.SessionKey)
	}
	log.Info("no metadata")
	return ""
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
	sizeService service.SizeService,
	productService service.ProductService,
	finalProductService service.FinalProductService,
) *ProductApiServiceServer {
	return &ProductApiServiceServer{
		categoryService:     categoryService,
		brandService:        brandService,
		sizeService:         sizeService,
		productService:      productService,
		finalProductService: finalProductService,
	}
}
