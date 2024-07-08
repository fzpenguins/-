package minio

import (
	"context"
	"grpc/config"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func InitMinIoClient() {
	var err error
	MinioClient, err = minio.New(config.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.SSL,
	})
	if err != nil {
		log.Fatalln("minio new err = ", err)
	}
	log.Println(1)
	err = MinioClient.MakeBucket(context.Background(), config.BucketName, minio.MakeBucketOptions{})
	if err != nil {
		log.Println("minio make bucket err = ", err)
	}
	log.Println(2)
}
