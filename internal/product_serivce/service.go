package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/config"
	"github.com/alexeyzer/product-api/internal/pkg/service"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"

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

func (s *ProductApiServiceServer) GetSessionIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		val := md.Get(config.Config.Auth.SessionKey)
		if len(val) > 0 {
			return val[0]
		}
	}
	return ""
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
