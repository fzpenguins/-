package rpc

import (
	"context"
	"grpc/proto"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitConvertRPC() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	// defer conn.Close()
	pictureConClient = proto.NewPictureConServiceClient(conn)
	//convertClient = p.NewClipServiceClient(conn)
}

func GetImageVector(ctx context.Context, req *proto.ImageRequest) (vector []float32, err error) {
	resp, err := pictureConClient.GetImageVector(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "rpc.GetImageVector failed")
	}
	return resp.Vector, nil
}
