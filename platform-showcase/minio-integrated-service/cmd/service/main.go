package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	useSSL = false

	minioEndpoint   string
	accessKeyID     string
	secretAccessKey string
)

func init() {
	_ = godotenv.Load(".platform/env/.env")

	minioEndpoint = os.Getenv("MINIO_ENDPOINT")
	accessKeyID = os.Getenv("MINIO_ROOT_USER")
	secretAccessKey = os.Getenv("MINIO_ROOT_PASSWORD")
}

func main() {
	ctx := context.Background()

	minioClient, err := minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient)

	bucketName := "minio-integrated-service"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	objectName := "object-directory/object.txt"
	filePath := "db/s3/minio-local/seed/object.txt"
	contentType := "text/plain"

	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}
