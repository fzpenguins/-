package main

import (
	"grpc/biz/rpc"
	"grpc/config"
	"grpc/dal"
	"log"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Init() {

	config.ReadConfig()
	dal.Init()

	log.Println("successfully running...")
	klog.SetLevel(klog.LevelDebug)
	rpc.Init()

}

func main() {
	Init()

}
