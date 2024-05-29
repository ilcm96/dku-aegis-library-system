package db

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"os"
)

var minioClient *minio.Client

func InitMinioClient() {
	var err error
	minioClient, err = minio.New(fmt.Sprintf("%s:%s", os.Getenv("MINIO_HOST"), os.Getenv("MINIO_PORT")), &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_KEY"), os.Getenv("MINIO_SECRET"), ""),
		Secure: false,
	})
	if err != nil {
		log.Panic(err)
	}
}

func MinioClient() *minio.Client {
	return minioClient
}
