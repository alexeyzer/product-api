package client

import (
	"context"
	pb "github.com/alexeyzer/product-api/pb/api/recognize/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecognizeAPIClient interface {
	RecognizePhoto(ctx context.Context, image []byte) (int64, error)
}

type recognizeAPIClient struct {
	classMap map[string]int64
	conn     *grpc.ClientConn
	client   pb.RecognizeApiServiceClient
}

func (r *recognizeAPIClient) RecognizePhoto(ctx context.Context, image []byte) (int64, error) {
	res, err := r.client.RecognizePhoto(ctx, &pb.RecognizePhotoRequest{
		Image: image,
	})
	if err != nil {
		return 0, err
	}
	category, ok := r.classMap[res.Category]
	if !ok {
		return 0, status.Errorf(codes.Internal, "unknown category: %s", res.Category)
	}

	return category, nil
}

func (r *recognizeAPIClient) classMapper() {

}

func NewRecognizeApiClient(address string) (RecognizeAPIClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewRecognizeApiServiceClient(conn)

	client := &recognizeAPIClient{
		classMap: map[string]int64{
			"Hoodie": 6,
			"Skirt":  7,
			"Tee":    5,
		},
		conn:   conn,
		client: c,
	}
	return client, nil
}
