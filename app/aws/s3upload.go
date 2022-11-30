package aws

import (
	util "PPOB_BACKEND/utils"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/echo/v4"
)

func UploadToS3(c echo.Context, filename string, src multipart.File) (string, error) {
	SECRET_KEY := util.GetEnv("AWS_S3_BUCKET_SECRET_KEY")
	KEY_ID := util.GetEnv("AWS_S3_BUCKET_KEY_ID")
	REGION := util.GetEnv("AWS_S3_REGION")

	configS3 := &aws.Config{
		Region:      aws.String(REGION),
		Credentials: credentials.NewStaticCredentials(KEY_ID, SECRET_KEY, ""),
	}
	s3Session := session.New(configS3)
	uploader := s3manager.NewUploader(s3Session)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("testingmycuans3"),
		Key:    aws.String("img/" + filename),
		Body:   src,
	})
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return result.Location, nil
}
