package main

import (
	"context"
	"encoding/json"
	"github.com/alexeyzer/product-api/internal/client"
	"github.com/alexeyzer/product-api/internal/pkg/service"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
	"time"

	"github.com/alexeyzer/product-api/config"
	"github.com/alexeyzer/product-api/internal/pkg/repository"
	"github.com/alexeyzer/product-api/internal/product_serivce"
	gw "github.com/alexeyzer/product-api/pb/api/product/v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var userAPIClient client.UserAPIClient

func serveSwagger(mux *http.ServeMux) {
	prefix := "/swagger/"
	sh := http.StripPrefix(prefix, http.FileServer(http.Dir("./swagger/")))
	mux.Handle(prefix, sh)
}

// look up session in cookie and pass sessionId in to context if it exists
func gatewayMetadataAnnotator(ctx context.Context, r *http.Request) metadata.MD {
	sessionID, ok := r.Cookie(config.Config.Auth.SessionKey)
	if ok == nil {
		md := metadata.Pairs(config.Config.Auth.SessionKey, sessionID.Value)
		ctx = metadata.AppendToOutgoingContext(ctx, config.Config.Auth.SessionKey, sessionID.Value)
		res, err := userAPIClient.SessionCheck(ctx)
		if err != nil {
			log.Info("failed to enrich metadata from user-api")
		} else {
			byte, err := json.Marshal(res)
			if err != nil {
				log.Info("failed to convert userinfo to byte")
			}
			md = metadata.Pairs(config.Config.Auth.UserInfoKey, string(byte))
		}

		return md
	}
	log.Println("No Cookie")
	return metadata.Pairs()
}

func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("origin"))
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		if (*r).Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func RunServer(ctx context.Context, productApiServiceServer *product_serivce.ProductApiServiceServer) error {

	grpcLis, err := net.Listen("tcp", ":"+config.Config.App.GrpcPort)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_logrus.UnaryServerInterceptor(log.WithContext(ctx).WithTime(time.Time{})),
		grpc_validator.UnaryServerInterceptor(),
		grpc_prometheus.UnaryServerInterceptor),
	))
	gw.RegisterProductApiServiceServer(grpcServer, productApiServiceServer)
	grpc_prometheus.Register(grpcServer)

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux(runtime.WithMetadata(gatewayMetadataAnnotator))
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/", corsMiddleware(gwmux))
	serveSwagger(mux)
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(),
	}
	err = gw.RegisterProductApiServiceHandlerFromEndpoint(ctx, gwmux, ":"+config.Config.App.GrpcPort, opts)
	if err != nil {
		return err
	}
	go func() {
		err = grpcServer.Serve(grpcLis)
		log.Fatal(err)
	}()
	log.Println("app started")
	err = http.ListenAndServeTLS(":"+config.Config.App.HttpPort, "./keys/server.crt", "./keys/server.key", mux)
	return err
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err := config.ReadConf("./config/config.yaml")
	if err != nil {
		log.Fatal("Failed to create config: ", err)
	}

	userAPIClient, err = client.NewUserApiClient(config.Config.GRPC.UserAPI)
	if err != nil {
		log.Fatal("Failed to connect to userAPI: ", err)
	}

	s3, err := client.NewS3Client(config.Config.S3.BucketName, config.Config.S3.ID, config.Config.S3.Key, config.Config.S3.Region, config.Config.S3.Endpoint)
	if err != nil {
		log.Fatal("Failed to connect to aws bucket: ", err)
	}

	dao, err := repository.NewDao()
	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
	}

	brandService := service.NewBrandService(dao, s3)
	categoryService := service.NewCategoryService(dao)
	colorService := service.NewColorService(dao)
	sizeService := service.NewSizeService(dao)
	_ = service.NewMediaService(dao, s3)
	productService := service.NewProductService(dao, s3)
	finalProductService := service.NewFinalProductService(dao)

	productApiServiceServer := product_serivce.NewProductApiServiceServer(categoryService, brandService, colorService, sizeService, productService, finalProductService)
	if err := RunServer(ctx, productApiServiceServer); err != nil {
		log.Fatal(err)
	}
}
