package client

import (
	"context"
	pb "github.com/alexeyzer/product-api/pb/api/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserAPIClient interface {
	SessionCheck(ctx context.Context) (*pb.SessionCheckResponse, error)
	GetUserInfoAboutProduct(ctx context.Context, req *pb.GetUserInfoAboutProductRequest) (*pb.GetUserInfoAboutProductResponse, error)
}

type userAPIClient struct {
	conn   *grpc.ClientConn
	client pb.UserApiServiceClient
}

func (c *userAPIClient) GetUserInfoAboutProduct(ctx context.Context, req *pb.GetUserInfoAboutProductRequest) (*pb.GetUserInfoAboutProductResponse, error) {
	resp, err := c.client.GetUserInfoAboutProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *userAPIClient) SessionCheck(ctx context.Context) (*pb.SessionCheckResponse, error) {
	res, err := c.client.SessionCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewUserApiClient(address string) (UserAPIClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewUserApiServiceClient(conn)

	client := &userAPIClient{
		conn:   conn,
		client: c,
	}
	return client, nil
}
