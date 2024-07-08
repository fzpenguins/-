package minio

import (
	"context"
	"grpc/config"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func UploadImage(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	objectName := "picture/" + uuid.Must(uuid.NewRandom()).String() + ext
	if filePath[0:4] == "/hom" {
		resp, err := os.Open(filePath)
		if err != nil {
			log.Println("failed to open the file")
			return "", err
		}
		_, err = MinioClient.PutObject(context.Background(), config.BucketName, objectName, resp, -1, minio.PutObjectOptions{})
		if err != nil {
			return "", err
		}

	} else {
		resp, err := http.Get(filePath)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		_, err = MinioClient.PutObject(context.Background(), config.BucketName, objectName, resp.Body, -1, minio.PutObjectOptions{})
		if err != nil {
			return "", err
		}
	}

	return objectName, nil
}
