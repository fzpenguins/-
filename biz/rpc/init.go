package rpc

import "grpc/proto"

var (
	pictureConClient proto.PictureConServiceClient
	ClientReady      = make(chan bool)
)

func Init() {
	InitConvertRPC()
	// ClientReady <- true
}
