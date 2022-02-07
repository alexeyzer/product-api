package product_serivce

import (
	"context"
	"github.com/alexeyzer/product-api/config"
	desc "github.com/alexeyzer/product-api/pb/api/product/v1"

	"google.golang.org/grpc/metadata"
)

type ProductApiServiceServer struct {
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

func NewProductApiServiceServer() *ProductApiServiceServer {
	return &ProductApiServiceServer{}
}
