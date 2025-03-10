package internal

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"os"
)

const YaclS3Region = "ru-central1"
const YaclS3Endpoint = "https://storage.yandexcloud.net/"

func Upload2S3(bucket string, key string, file io.Reader) (*manager.UploadOutput, error) {
	accessKey, exist := os.LookupEnv("AWS_ACCESS_KEY_ID")
	if !exist {
		log.Fatalln("Ошибка загрузки в S3: AWS_ACCESS_KEY_ID нет в ENV")
	}
	secretKey, exist := os.LookupEnv("AWS_SECRET_ACCESS_KEY")
	if !exist {
		log.Fatalln("Ошибка загрузки в S3: AWS_SECRET_ACCESS_KEY нет в ENV")
	}

	cfg := aws.Config{
		Region:       YaclS3Region,
		BaseEndpoint: aws.String(YaclS3Endpoint),
		Credentials:  aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	return uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	})
}
