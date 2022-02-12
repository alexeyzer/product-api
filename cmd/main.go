package main

import (
	"context"
	"github.com/alexeyzer/product-api/internal/client"
	"github.com/alexeyzer/product-api/internal/pkg/service"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func authFunc(ctx context.Context) (context.Context, error) {
	log.Info("Auth func")
	if !config.Config.Auth.Working {
		return ctx, nil
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		val := md.Get(config.Config.Auth.SessionKey)
		if len(val) > 0 {
			ctx := metadata.AppendToOutgoingContext(ctx, config.Config.Auth.SessionKey, val[0])
			res, err := userAPIClient.SessionCheck(ctx)
			if err != nil {
				return nil, status.Error(codes.Unauthenticated, err.Error())
			}
			log.Info(res)
		} else {
			return nil, status.Error(codes.Unauthenticated, "SessionID doesn't exist")
		}
	} else {
		return nil, status.Error(codes.Unauthenticated, "SessionID doesn't exist")
	}
	return ctx, nil
}

func serveSwagger(mux *http.ServeMux) {
	prefix := "/swagger/"
	sh := http.StripPrefix(prefix, http.FileServer(http.Dir("./swagger/")))
	mux.Handle(prefix, sh)
}

// look up session and pass sessionId in to context if it exists
func gatewayMetadataAnnotator(_ context.Context, r *http.Request) metadata.MD {
	SessionID, ok := r.Cookie(config.Config.Auth.SessionKey)
	if ok == nil {
		log.Println(SessionID, ok)
		return metadata.Pairs(config.Config.Auth.SessionKey, SessionID.Value)
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
		grpc_auth.UnaryServerInterceptor(authFunc)),
	))
	gw.RegisterProductApiServiceServer(grpcServer, productApiServiceServer)

	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux(runtime.WithMetadata(gatewayMetadataAnnotator))
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

	dao, err := repository.NewDao()
	if err != nil {
		log.Fatal("Failed to connect to db: ", err)
	}

	categoryService := service.NewCategoryService(dao)

	productApiServiceServer := product_serivce.NewProductApiServiceServer(categoryService)
	if err := RunServer(ctx, productApiServiceServer); err != nil {
		log.Fatal(err)
	}
}
