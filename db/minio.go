package db

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var minioClient *minio.Client

func InitMinioClient() {
	var err error
	minioClient, err = minio.New("localhost:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("QePrC0jLCiYhy1etE1Pj", "e5ykM4axNx1kDBGIQp91SsWfPVbPo4d3vZq6wzcJ", ""),
		Secure: false,
	})
	if err != nil {
		log.Panic(err)
	}
}

func MinioClient() *minio.Client {
	return minioClient
}
