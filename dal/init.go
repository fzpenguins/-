package dal

import (
	"grpc/dal/db"
	"grpc/dal/milvus"
	"grpc/dal/minio"
	"log"
)

func Init() {
	db.Init()
	log.Println("successfully connect to database")
	milvus.Init()
	log.Println("successfully connect to milvus")
	minio.InitMinIoClient()
	log.Println("successfully connect to minio")
}
